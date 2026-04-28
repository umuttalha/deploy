package config

import (
	"fmt"
	"slices"
)

// Validate returns an error if the Stack is incomplete or inconsistent.
func Validate(s Stack) error {
	if s.Name == "" {
		return fmt.Errorf("stack name is required")
	}
	if s.Region == "" {
		return fmt.Errorf("region is required")
	}
	if !slices.Contains(s.Storage.Values(), string(s.Storage)) {
		return fmt.Errorf("invalid storage %q (want one of %v)", s.Storage, s.Storage.Values())
	}
	if s.InstanceType == "" {
		return fmt.Errorf("instance type is required")
	}
	if !slices.Contains(s.Image.Registry.Values(), string(s.Image.Registry)) {
		return fmt.Errorf("invalid image registry %q (want one of %v)", s.Image.Registry, s.Image.Registry.Values())
	}
	if s.Image.Repo == "" {
		return fmt.Errorf("image repo is required")
	}
	if s.Image.Tag == "" {
		return fmt.Errorf("image tag is required")
	}
	if !slices.Contains(s.DNS.Values(), string(s.DNS)) {
		return fmt.Errorf("invalid dns mode %q (want one of %v)", s.DNS, s.DNS.Values())
	}
	if !slices.Contains(s.Network.Values(), string(s.Network)) {
		return fmt.Errorf("invalid network %q (want one of %v)", s.Network, s.Network.Values())
	}
	return nil
}
