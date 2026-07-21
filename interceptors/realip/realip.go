package realip

import (
	"context"
	"net"
	"net/netip"

	"google.golang.org/grpc"
)

const (
	XRealIp       = "X-Real-IP"
	XForwardedFor = "X-Forwarded-For"
	TrueClientIp  = "True-Client-IP"
)

var noIP = netip.Addr{}

type realipKey struct{}

func FromContext(ctx context.Context) (netip.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), false
}

func remotePeer(ctx context.Context) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func ipInNets(ip netip.Addr, nets []netip.Prefix) bool { _ = "STUB: not implemented"; return false }

func getHeader(ctx context.Context, key string) string { _ = "STUB: not implemented"; return "" }

func ipFromXForwardedFoR(trustedProxies []netip.Prefix, ips []string, idx int) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

func ipFromHeaders(ctx context.Context, headers []string, trustedProxies []netip.Prefix, trustedProxyCnt uint) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

func getRemoteIP(ctx context.Context, trustedPeers, trustedProxies []netip.Prefix, headers []string, proxyCnt uint) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *serverStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func UnaryServerInterceptor(trustedPeers []netip.Prefix, headers []string) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptor(trustedPeers []netip.Prefix, headers []string) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func UnaryServerInterceptorOpts(opts ...Option) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func StreamServerInterceptorOpts(opts ...Option) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
