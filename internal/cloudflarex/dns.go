package cloudflarex

import (
	"context"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// CreateDNS provisions Cloudflare DNS records honoring s.DNS (Proxied vs DNSOnly).
// Cloudflare DNS records support a Tags field; the standard tag set is applied there.
func CreateDNS(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name)
	// TODO: c.API.DNS.Records.New(ctx, ...) with proxied=(s.DNS == config.DNSProxied).
	return nil
}

// DestroyDNS deletes DNS records tagged stack=<stackName>.
func DestroyDNS(ctx context.Context, c *Clients, stackName string) error {
	// TODO: list records, filter by tag, delete.
	return nil
}

// FindDNS returns DNS records tagged stack=<stackName>.
func FindDNS(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.API.DNS.Records.List(ctx, ...) filtered by tag.
	return nil, nil
}
