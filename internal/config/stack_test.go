package config

import (
	"reflect"
	"testing"
)

func TestStorageTypeValues(t *testing.T) {
	got := StorageS3.Values()
	want := []string{"s3", "r2"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("StorageType.Values() = %v, want %v", got, want)
	}
}

func TestDNSModeValues(t *testing.T) {
	got := DNSProxied.Values()
	want := []string{"proxied", "dns-only"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DNSMode.Values() = %v, want %v", got, want)
	}
}

func TestNetworkModeValues(t *testing.T) {
	got := NetIPv4.Values()
	want := []string{"ipv4", "ipv6", "dual"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NetworkMode.Values() = %v, want %v", got, want)
	}
}

func TestImageRegistryValues(t *testing.T) {
	got := RegistryGHCR.Values()
	want := []string{"ghcr", "ecr"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ImageRegistry.Values() = %v, want %v", got, want)
	}
}
