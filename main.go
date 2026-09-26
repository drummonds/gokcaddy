// Command gokcaddy is Caddy built with the Route 53 DNS provider, for the
// gokrazy instances under ~/minor/gok_local. The stock
// github.com/caddyserver/caddy/v2/cmd/caddy package has no DNS providers,
// and gokrazy builds packages as-is, so a custom main is the only way to get
// the DNS-01 challenge (wildcard certificates for LAN hosts) onto the device.
//
// This is what xcaddy would generate; keeping it as a real module means gok
// can fetch and update it like any other package.
package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	_ "github.com/caddy-dns/route53"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

func main() {
	caddycmd.Main()
}
