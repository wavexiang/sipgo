package sip

import (
	"context"
	"net"
	"strconv"
)

var (

	// IdleConnection will keep connections idle even after transaction terminate
	// -1 	- single response or request will close
	// 0 	- close connection immediatelly after transaction terminate
	// 1 	- keep connection idle after transaction termination
	IdleConnection int = 1
)

const (
	MTU uint = 1500

	DefaultHost     = "127.0.0.1"
	DefaultProtocol = "UDP"

	DefaultUdpPort int = 5060
	DefaultTcpPort int = 5060
	DefaultTlsPort int = 5061
	DefaultWsPort  int = 80
	DefaultWssPort int = 443
	// Transport for different sip messages. GO uses lowercase, but for message parsing, we should
	// use this constants for setting message Transport
	TransportUDP = "UDP"
	TransportTCP = "TCP"
	TransportTLS = "TLS"
	TransportWS  = "WS"
	TransportWSS = "WSS"

	transportBufferSize uint16 = 65535

	// TransportFixedLengthMessage sets message size limit for parsing and avoids stream parsing
	TransportFixedLengthMessage uint16 = 0
)

// Protocol implements network specific features.
type Transport interface {
	Network() string

	// GetConnection returns connection from transport
	// addr must be resolved to IP:port
	GetConnection(addr string) (Connection, error)
	CreateConnection(ctx context.Context, laddr Addr, raddr Addr, handler MessageHandler) (Connection, error)
	String() string
	Close() error
}

// DefaultPort returns transport default port by network.
func DefaultPort(transport string) int {
	switch ASCIIToLower(transport) {
	case "tls":
		return DefaultTlsPort
	case "tcp":
		return DefaultTcpPort
	case "udp":
		return DefaultUdpPort
	case "ws":
		return DefaultWsPort
	case "wss":
		return DefaultWssPort
	default:
		return DefaultTcpPort
	}
}

type Addr struct {
	IP       net.IP // Must be in IP format
	Port     int
	Hostname string // Original hostname before resolved to IP
}

func (a *Addr) String() string {
	if a.IP == nil {
		return net.JoinHostPort(a.Hostname, strconv.Itoa(a.Port))
	}

	return net.JoinHostPort(a.IP.String(), strconv.Itoa(a.Port))
}

func ParseAddr(addr string) (host string, port int, err error) {
	host, pstr, err := net.SplitHostPort(addr)
	if err != nil {
		return host, port, err
	}

	// In case we are dealing with some named ports this should be called
	// net.LookupPort(network)

	port, err = strconv.Atoi(pstr)
	return host, port, err
}
