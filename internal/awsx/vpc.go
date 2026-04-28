package awsx

import (
	"context"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/tagging"
)

// CreateVPC provisions a VPC honoring s.Network (IPv4/IPv6/dual) and s.AllowNAT.
// When AllowNAT is false, no NAT Gateway is created — private subnets must use
// VPC endpoints or be design-eliminated.
func CreateVPC(ctx context.Context, c *Clients, s config.Stack) error {
	_ = tagging.StackTags(s.Name)
	// TODO: c.EC2.CreateVpc, CreateSubnet, CreateInternetGateway.
	// TODO: if s.AllowNAT { c.EC2.CreateNatGateway(...) }.
	// TODO: dual-stack: assign Ipv6CidrBlock and AssignIpv6AddressOnCreation.
	return nil
}

// DestroyVPC tears down VPC + subnets + IGW (+ NAT if any) tagged stack=<stackName>.
func DestroyVPC(ctx context.Context, c *Clients, stackName string) error {
	// TODO: reverse-dependency teardown using c.Tagging to find resources.
	return nil
}

// FindVPC returns VPC-related resources tagged stack=<stackName>.
func FindVPC(ctx context.Context, c *Clients, stackName string) ([]Resource, error) {
	// TODO: c.Tagging.GetResources with ResourceTypeFilters: []string{"ec2:vpc","ec2:subnet"}.
	return nil, nil
}
