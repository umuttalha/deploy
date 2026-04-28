package config

// StorageType selects where bucket-style storage lives.
type StorageType string

const (
	StorageS3 StorageType = "s3"
	StorageR2 StorageType = "r2"
)

func (StorageType) Values() []string { return []string{"s3", "r2"} }

// DNSMode selects whether Cloudflare proxies traffic through its edge.
type DNSMode string

const (
	DNSProxied DNSMode = "proxied"
	DNSOnly    DNSMode = "dns-only"
)

func (DNSMode) Values() []string { return []string{"proxied", "dns-only"} }

// NetworkMode selects the IP family for VPC subnets.
type NetworkMode string

const (
	NetIPv4      NetworkMode = "ipv4"
	NetIPv6      NetworkMode = "ipv6"
	NetDualStack NetworkMode = "dual"
)

func (NetworkMode) Values() []string { return []string{"ipv4", "ipv6", "dual"} }

// ImageRegistry routes container image pulls.
type ImageRegistry string

const (
	RegistryGHCR ImageRegistry = "ghcr"
	RegistryECR  ImageRegistry = "ecr"
)

func (ImageRegistry) Values() []string { return []string{"ghcr", "ecr"} }

// ImageRef points at a specific container image.
type ImageRef struct {
	Registry ImageRegistry
	Repo     string
	Tag      string
}

// Stack is the single source of truth for user-selected options.
// Each field maps to one cobra flag and one huh prompt fallback.
type Stack struct {
	Name         string
	Region       string
	Storage      StorageType
	InstanceType string
	Image        ImageRef
	DNS          DNSMode
	Network      NetworkMode
	AllowNAT     bool
}
