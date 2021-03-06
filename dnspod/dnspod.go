// Package dnspod adapts the lego Dnspod DNS
// provider for Caddy. Importing this package plugs it in.
package dnspod

import (
	"errors"

	"github.com/icasei/caddy/caddytls"
	"github.com/icasei/lego/providers/dns/dnspod"
)

func init() {
	caddytls.RegisterDNSProvider("dnspod", NewDNSProvider)
}

// NewDNSProvider returns a new dnspod DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = access token (API key)
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dnspod.NewDNSProvider()
	case 1:
		return dnspod.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
