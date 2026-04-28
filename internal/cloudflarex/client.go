// Package cloudflarex wires the Cloudflare SDK and provides the same
// Create/Destroy/Find contract as awsx for parity.
package cloudflarex

import (
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// Clients bundles the Cloudflare API client.
type Clients struct {
	API *cloudflare.Client
}

// Resource mirrors awsx.Resource for parity in `deploy list` output.
type Resource struct {
	Kind string
	ID   string
}

// New constructs the Cloudflare client. Token is required.
func New(token string) *Clients {
	return &Clients{
		API: cloudflare.NewClient(option.WithAPIToken(token)),
	}
}
