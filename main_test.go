package main

import (
	"testing"

	"github.com/caddyserver/caddy/v2"
)

// The whole point of this module is Caddy plus the Route 53 DNS provider.
// If the provider ever drops out of the build, DNS-01 issuance on the
// gokrazy instances silently stops at the next renewal.
func TestRoute53ProviderIsCompiledIn(t *testing.T) {
	if _, err := caddy.GetModule("dns.providers.route53"); err != nil {
		t.Fatalf("dns.providers.route53 not registered: %v", err)
	}
}
