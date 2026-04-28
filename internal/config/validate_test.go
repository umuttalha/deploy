package config

import (
	"strings"
	"testing"
)

func validStack() Stack {
	return Stack{
		Name:         "demo",
		Region:       "us-east-1",
		Storage:      StorageS3,
		InstanceType: "t3.micro",
		Image:        ImageRef{Registry: RegistryGHCR, Repo: "umuttalha/app", Tag: "latest"},
		DNS:          DNSProxied,
		Network:      NetIPv4,
		AllowNAT:     false,
	}
}

func TestValidate_OK(t *testing.T) {
	if err := Validate(validStack()); err != nil {
		t.Fatalf("Validate(valid) = %v, want nil", err)
	}
}

func TestValidate_MissingName(t *testing.T) {
	s := validStack()
	s.Name = ""
	err := Validate(s)
	if err == nil || !strings.Contains(err.Error(), "name") {
		t.Fatalf("Validate(no name) = %v, want error mentioning name", err)
	}
}

func TestValidate_BadStorage(t *testing.T) {
	s := validStack()
	s.Storage = "bogus"
	err := Validate(s)
	if err == nil || !strings.Contains(err.Error(), "storage") {
		t.Fatalf("Validate(bad storage) = %v, want error mentioning storage", err)
	}
}

func TestValidate_BadNetwork(t *testing.T) {
	s := validStack()
	s.Network = "ipv7"
	err := Validate(s)
	if err == nil || !strings.Contains(err.Error(), "network") {
		t.Fatalf("Validate(bad network) = %v, want error mentioning network", err)
	}
}

func TestValidate_MissingImageRepo(t *testing.T) {
	s := validStack()
	s.Image.Repo = ""
	err := Validate(s)
	if err == nil || !strings.Contains(err.Error(), "image") {
		t.Fatalf("Validate(no image repo) = %v, want error mentioning image", err)
	}
}
