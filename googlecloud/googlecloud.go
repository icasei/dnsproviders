// Package googlecloud adapts the lego Google Cloud DNS
// provider for Caddy. Importing this package plugs it in.
package googlecloud

import (
	"errors"

	"github.com/icasei/caddy/caddytls"
	"github.com/icasei/lego/providers/dns/gcloud"
)

func init() {
	caddytls.RegisterDNSProvider("googlecloud", NewDNSProvider)
}

// NewDNSProvider returns a new Google Cloud DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = project
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return gcloud.NewDNSProvider()
	case 1:
		return gcloud.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
