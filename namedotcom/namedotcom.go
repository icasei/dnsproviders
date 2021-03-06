// Package namedotcom adapts the lego Name.com DNS provider
// for Caddy. Importing this package plugs it in.
package namedotcom

import (
	"errors"

	"github.com/icasei/caddy/caddytls"
	"github.com/icasei/lego/providers/dns/namedotcom"
)

func init() {
	caddytls.RegisterDNSProvider("namedotcom", NewDNSProvider)
}

// NewDNSProvider returns a new Name.com DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = username
//         credentials[1] = API token
// len(3): credentials[0] = username
//         credentials[1] = API token
//         credentials[2] = Server
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return namedotcom.NewDNSProvider()
	case 2:
		return namedotcom.NewDNSProviderCredentials(credentials[0], credentials[1], "")
	case 3:
		return namedotcom.NewDNSProviderCredentials(credentials[0], credentials[1], credentials[2])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
