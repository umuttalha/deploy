// Package ghcr wires GitHub Container Registry interactions. GHCR lives on
// GitHub, not Cloudflare or AWS, so it has its own package. The exported
// Create/Destroy/Find shape mirrors awsx and cloudflarex for parity.
package ghcr

import (
	"context"

	"github.com/google/go-github/v66/github"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// Clients bundles the GitHub client.
type Clients struct {
	GH *github.Client
}

// Resource mirrors awsx.Resource for parity in `deploy list` output.
type Resource struct {
	Kind string
	ID   string
}

// New constructs the GitHub client with the supplied token. Token may be empty
// for unauthenticated reads of public images, but writes require a token.
func New(token string) *Clients {
	c := github.NewClient(nil)
	if token != "" {
		c = c.WithAuthToken(token)
	}
	return &Clients{GH: c}
}

// CreatePackage ensures the GHCR package referenced by s.Image exists and is
// tagged via repository labels (GHCR has no native tag API).
func CreatePackage(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name)
	// TODO: c.GH.Organizations.GetPackage(ctx, ...) or PUT label on the source repo.
	return nil
}

// DestroyPackage deletes GHCR package versions associated with stackName.
func DestroyPackage(ctx context.Context, c *Clients, stackName string) error {
	// TODO: list versions, filter by label, delete.
	return nil
}

// FindPackage returns GHCR packages associated with stackName.
func FindPackage(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.GH.Organizations.ListPackages(ctx, ...) filtered by label.
	return nil, nil
}
