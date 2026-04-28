package cloudflarex

import (
	"context"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// CreateR2 provisions a Cloudflare R2 bucket. R2 buckets do not natively
// support tags via the public API today, so the standard tag set is
// recorded as a name prefix or a sidecar metadata object.
func CreateR2(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name)
	// TODO: c.API.R2.Buckets.New(ctx, ...)
	return nil
}

// DestroyR2 deletes R2 buckets associated with stackName.
func DestroyR2(ctx context.Context, c *Clients, stackName string) error {
	// TODO: list buckets, filter by name prefix matching stackName, delete.
	return nil
}

// FindR2 returns R2 buckets associated with stackName.
func FindR2(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.API.R2.Buckets.List(ctx, ...) and filter.
	return nil, nil
}
