package awsx

import (
	"context"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// ----------------------------------------------------------------------------
// s3.go is the REFERENCE TEMPLATE for every provisioner file in awsx/,
// cloudflarex/, and ghcr/. The contract is:
//
//   - Create applies tagging.StackTags(s.Name) to every resource it creates.
//   - Destroy looks up resources by stack tag, then deletes them.
//   - Find returns the live list, used by `deploy list`.
//
// Every other file in awsx/ (ec2.go, vpc.go, ecr.go) and the parallel files
// in cloudflarex/ and ghcr/ mirror this shape. Bodies are stubbed; fill them
// in one resource at a time.
// ----------------------------------------------------------------------------

// Create provisions the S3 bucket described by s and tags it with the
// standard managed-by + stack tags. Idempotent: re-running with the same
// stack name should not create duplicate buckets.
func Create(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name) // standard tag set goes on every API call
	// TODO: c.S3.CreateBucket(ctx, &s3.CreateBucketInput{...})
	// TODO: c.S3.PutBucketTagging(ctx, ...)
	return nil
}

// Destroy deletes every S3 bucket tagged stack=<stackName>.
func Destroy(ctx context.Context, c *Clients, stackName string) error {
	// TODO: list buckets, filter by tag via c.Tagging, c.S3.DeleteBucket(ctx, ...)
	return nil
}

// Find returns S3 buckets tagged stack=<stackName>.
func Find(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.Tagging.GetResources(ctx, &resourcegroupstaggingapi.GetResourcesInput{
	//   TagFilters: []types.TagFilter{
	//     {Key: aws.String(tagging.ManagedByKey), Values: []string{tagging.ManagedByValue}},
	//     {Key: aws.String(tagging.StackKey),     Values: []string{stackName}},
	//   },
	//   ResourceTypeFilters: []string{"s3"},
	// })
	return nil, nil
}
