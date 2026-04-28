package awsx

import (
	"context"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// CreateECR creates an ECR repository sized for s.Image.Repo.
func CreateECR(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name)
	// TODO: c.ECR.CreateRepository(ctx, &ecr.CreateRepositoryInput{RepositoryName: ..., Tags: ...})
	return nil
}

// DestroyECR deletes ECR repos tagged stack=<stackName>.
func DestroyECR(ctx context.Context, c *Clients, stackName string) error {
	// TODO: c.ECR.DescribeRepositories filtered by tag, then DeleteRepository.
	return nil
}

// FindECR returns ECR repos tagged stack=<stackName>.
func FindECR(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.Tagging.GetResources with ResourceTypeFilters: []string{"ecr:repository"}.
	return nil, nil
}
