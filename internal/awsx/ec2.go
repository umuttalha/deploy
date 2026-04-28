package awsx

import (
	"context"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// CreateEC2 launches an EC2 instance of type s.InstanceType in s.Region,
// tagged with the standard tag set.
func CreateEC2(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name)
	// TODO: c.EC2.RunInstances(ctx, &ec2.RunInstancesInput{InstanceType: types.InstanceType(s.InstanceType), ...})
	return nil
}

// DestroyEC2 terminates EC2 instances tagged stack=<stackName>.
func DestroyEC2(ctx context.Context, c *Clients, stackName string) error {
	// TODO: c.EC2.DescribeInstances with tag filter, then TerminateInstances.
	return nil
}

// FindEC2 returns EC2 instances tagged stack=<stackName>.
func FindEC2(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.Tagging.GetResources with ResourceTypeFilters: []string{"ec2:instance"}.
	return nil, nil
}
