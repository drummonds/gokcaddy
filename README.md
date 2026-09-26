# gokcaddy

Caddy built with the Route 53 DNS provider, for gokrazy instances.

gokrazy builds Go packages as-is and the stock
`github.com/caddyserver/caddy/v2/cmd/caddy` has no DNS providers, so this
one-file `main` is what `xcaddy build --with github.com/caddy-dns/route53`
would produce, kept as a real module so `gok` can fetch and update it.

Use it in a gokrazy `config.json` in place of the stock package:

```json
"Packages": ["git.bytestone.uk/hum3/gokcaddy"],
"PackageConfig": {
  "git.bytestone.uk/hum3/gokcaddy": {
    "CommandLineFlags": ["run", "--config", "/etc/caddy/Caddyfile"],
    "Environment": ["AWS_ACCESS_KEY_ID=…", "AWS_SECRET_ACCESS_KEY=…"],
    "ExtraFilePaths": {"/etc/caddy/Caddyfile": "Caddyfile"}
  }
}
```

and in the Caddyfile `tls { dns route53 }`. gokrazy gives the service
`HOME=/perm/home/gokcaddy`, so certificates persist across updates.

## Links

- Source: https://git.bytestone.uk/hum3/gokcaddy
- Mirror: https://github.com/drummonds/gokcaddy
- Provider: https://github.com/caddy-dns/route53
