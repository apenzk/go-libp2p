package quicreuse

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"sync"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/quic-go/quic-go"
	quiclogging "github.com/quic-go/quic-go/logging"
)

type ConnManager struct {
	reuseUDP4       *reuse
	reuseUDP6       *reuse
	enableDraft29   bool
	enableReuseport bool
	enableMetrics   bool

	serverConfig *quic.Config
	clientConfig *quic.Config

	connsMu sync.Mutex
	conns   map[string]connListenerEntry

	srk           quic.StatelessResetKey
	metricsTracer *metricsTracer
}

type connListenerEntry struct {
	refCount int
	ln       *connListener
}

func NewConnManager(statelessResetKey quic.StatelessResetKey, opts ...Option) (*ConnManager, error) {
	cm := &ConnManager{
		enableReuseport: true,
		enableDraft29:   true,
		conns:           make(map[string]connListenerEntry),
		srk:             statelessResetKey,
	}
	for _, o := range opts {
		if err := o(cm); err != nil {
			return nil, err
		}
	}

	quicConf := quicConfig.Clone()

	var tracers []quiclogging.Tracer

	if cm.enableMetrics {
		cm.metricsTracer = newMetricsTracer()
	}
	if len(tracers) > 0 {
		quicConf.Tracer = func(ctx context.Context, p quiclogging.Perspective, ci quic.ConnectionID) quiclogging.ConnectionTracer {
			tracers := make([]quiclogging.ConnectionTracer, 0, 2)
			if qlogTracerDir != "" {
				tracers = append(tracers, qloggerForDir(qlogTracerDir, p, ci))
			}
			if cm.metricsTracer != nil {
				tracers = append(tracers, cm.metricsTracer.TracerForConnection(ctx, p, ci))
			}
			return quiclogging.NewMultiplexedConnectionTracer(tracers...)
		}
	}
	serverConfig := quicConf.Clone()
	if !cm.enableDraft29 {
		serverConfig.Versions = []quic.VersionNumber{quic.Version1}
	}

	cm.clientConfig = quicConf
	cm.serverConfig = serverConfig
	if cm.enableReuseport {
		cm.reuseUDP4 = newReuse(&statelessResetKey, cm.metricsTracer)
		cm.reuseUDP6 = newReuse(&statelessResetKey, cm.metricsTracer)
	}
	return cm, nil
}

func (c *ConnManager) getReuse(network string) (*reuse, error) {
	switch network {
	case "udp4":
		return c.reuseUDP4, nil
	case "udp6":
		return c.reuseUDP6, nil
	default:
		return nil, errors.New("invalid network: must be either udp4 or udp6")
	}
}

func (c *ConnManager) ListenQUIC(addr ma.Multiaddr, tlsConf *tls.Config, allowWindowIncrease func(conn quic.Connection, delta uint64) bool) (Listener, error) {
	if !c.enableDraft29 {
		if _, err := addr.ValueForProtocol(ma.P_QUIC); err == nil {
			return nil, errors.New("can't listen on `/quic` multiaddr (QUIC draft 29 version) when draft 29 support is disabled")
		}
	}

	netw, host, err := manet.DialArgs(addr)
	if err != nil {
		return nil, err
	}
	laddr, err := net.ResolveUDPAddr(netw, host)
	if err != nil {
		return nil, err
	}

	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	key := laddr.String()
	entry, ok := c.conns[key]
	if !ok {
		conn, err := c.transportForListen(netw, laddr)
		if err != nil {
			return nil, err
		}
		ln, err := newConnListener(conn, c.serverConfig, c.enableDraft29)
		if err != nil {
			return nil, err
		}
		key = conn.LocalAddr().String()
		entry = connListenerEntry{ln: ln}
	}
	l, err := entry.ln.Add(tlsConf, allowWindowIncrease, func() { c.onListenerClosed(key) })
	if err != nil {
		if entry.refCount <= 0 {
			entry.ln.Close()
		}
		return nil, err
	}
	entry.refCount++
	c.conns[key] = entry
	return l, nil
}

func (c *ConnManager) onListenerClosed(key string) {
	c.connsMu.Lock()
	defer c.connsMu.Unlock()

	entry := c.conns[key]
	entry.refCount = entry.refCount - 1
	if entry.refCount <= 0 {
		delete(c.conns, key)
		entry.ln.Close()
	} else {
		c.conns[key] = entry
	}
}

func (c *ConnManager) transportForListen(network string, laddr *net.UDPAddr) (refCountedQuicTransport, error) {
	if c.enableReuseport {
		reuse, err := c.getReuse(network)
		if err != nil {
			return nil, err
		}
		return reuse.TransportForListen(network, laddr)
	}

	conn, err := listenAndOptimize(network, laddr)
	if err != nil {
		return nil, err
	}
	return &singleOwnerTransport{tr: quic.Transport{Conn: conn, StatelessResetKey: &c.srk}, packetConn: conn}, nil
}

func (c *ConnManager) DialQUIC(ctx context.Context, raddr ma.Multiaddr, tlsConf *tls.Config, allowWindowIncrease func(conn quic.Connection, delta uint64) bool) (quic.Connection, error) {
	naddr, v, err := FromQuicMultiaddr(raddr)
	if err != nil {
		return nil, err
	}
	netw, _, err := manet.DialArgs(raddr)
	if err != nil {
		return nil, err
	}

	quicConf := c.clientConfig.Clone()
	quicConf.AllowConnectionWindowIncrease = allowWindowIncrease

	if v == quic.Version1 {
		// The endpoint has explicit support for QUIC v1, so we'll only use that version.
		quicConf.Versions = []quic.VersionNumber{quic.Version1}
	} else if v == quic.VersionDraft29 {
		quicConf.Versions = []quic.VersionNumber{quic.VersionDraft29}
	} else {
		return nil, errors.New("unknown QUIC version")
	}

	tr, err := c.TransportForDial(netw, naddr)
	if err != nil {
		return nil, err
	}
	conn, err := tr.Transport().Dial(ctx, naddr, tlsConf, quicConf)
	if err != nil {
		tr.DecreaseCount()
		return nil, err
	}
	return conn, nil
}

func (c *ConnManager) TransportForDial(network string, raddr *net.UDPAddr) (refCountedQuicTransport, error) {
	if c.enableReuseport {
		reuse, err := c.getReuse(network)
		if err != nil {
			return nil, err
		}
		return reuse.TransportForDial(network, raddr)
	}

	var laddr *net.UDPAddr
	switch network {
	case "udp4":
		laddr = &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	case "udp6":
		laddr = &net.UDPAddr{IP: net.IPv6zero, Port: 0}
	}
	conn, err := listenAndOptimize(network, laddr)
	if err != nil {
		return nil, err
	}
	return &singleOwnerTransport{tr: quic.Transport{Conn: conn, StatelessResetKey: &c.srk}, packetConn: conn}, nil
}

func (c *ConnManager) Protocols() []int {
	if c.enableDraft29 {
		return []int{ma.P_QUIC, ma.P_QUIC_V1}
	}
	return []int{ma.P_QUIC_V1}
}

func (c *ConnManager) Close() error {
	if !c.enableReuseport {
		return nil
	}
	if err := c.reuseUDP6.Close(); err != nil {
		return err
	}
	return c.reuseUDP4.Close()
}

// listenAndOptimize same as net.ListenUDP, but also calls quic.OptimizeConn
func listenAndOptimize(network string, laddr *net.UDPAddr) (net.PacketConn, error) {
	conn, err := net.ListenUDP(network, laddr)
	if err != nil {
		return nil, err
	}
	return quic.OptimizeConn(conn)
}
