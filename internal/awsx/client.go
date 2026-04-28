// Package awsx wires AWS SDK v2 service clients used by the deploy CLI.
// Each resource lives in its own file (s3.go, ec2.go, vpc.go, ecr.go) and
// exposes the same Create/Destroy/Find contract.
package awsx

import (
	"context"
	"fmt"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Clients bundles every AWS service client this CLI uses.
type Clients struct {
	Cfg     awssdk.Config
	S3      *s3.Client
	EC2     *ec2.Client
	ECR     *ecr.Client
	Tagging *resourcegroupstaggingapi.Client
}

// Resource is a thin descriptor returned from Find.
type Resource struct {
	Kind string // "s3-bucket", "ec2-instance", ...
	ID   string // ARN or service-specific identifier
}

// New builds a Clients with default credentials chain. Region and profile are optional.
func New(ctx context.Context, region, profile string) (*Clients, error) {
	opts := []func(*config.LoadOptions) error{}
	if region != "" {
		opts = append(opts, config.WithRegion(region))
	}
	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}
	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("aws: load config: %w", err)
	}
	return &Clients{
		Cfg:     cfg,
		S3:      s3.NewFromConfig(cfg),
		EC2:     ec2.NewFromConfig(cfg),
		ECR:     ecr.NewFromConfig(cfg),
		Tagging: resourcegroupstaggingapi.NewFromConfig(cfg),
	}, nil
}
