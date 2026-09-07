package realip

import "net/netip"

type options struct {
	trustedPeers []netip.Prefix

	trustedProxies []netip.Prefix

	trustedProxiesCount uint

	headers []string
}

type Option func(*options)

func evaluateOpts(opts []Option) *options { _ = "STUB: not implemented"; return nil }

func WithTrustedPeers(peers []netip.Prefix) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTrustedProxies(proxies []netip.Prefix) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTrustedProxiesCount(count uint) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeaders(headers []string) Option { _ = "STUB: not implemented"; return *new(Option) }
