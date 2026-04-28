package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// Prompt fills any zero-valued fields on s by asking the user via a huh form.
// Fields already populated (e.g. by cobra flags) are skipped.
func Prompt(s *Stack) error {
	var groups []*huh.Group

	if s.Name == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewInput().Title("Stack name").Value(&s.Name),
		))
	}
	if s.Region == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewInput().Title("AWS region").Placeholder("us-east-1").Value(&s.Region),
		))
	}
	if s.Storage == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[StorageType]().
				Title("Storage").
				Options(
					huh.NewOption("Amazon S3", StorageS3),
					huh.NewOption("Cloudflare R2", StorageR2),
				).
				Value(&s.Storage),
		))
	}
	if s.InstanceType == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewInput().Title("Instance type").Placeholder("t3.micro").Value(&s.InstanceType),
		))
	}
	if s.Image.Registry == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[ImageRegistry]().
				Title("Image registry").
				Options(
					huh.NewOption("GitHub Container Registry", RegistryGHCR),
					huh.NewOption("Amazon ECR", RegistryECR),
				).
				Value(&s.Image.Registry),
		))
	}
	if s.Image.Repo == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewInput().Title("Image repo").Placeholder("owner/app").Value(&s.Image.Repo),
		))
	}
	if s.Image.Tag == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewInput().Title("Image tag").Placeholder("latest").Value(&s.Image.Tag),
		))
	}
	if s.DNS == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[DNSMode]().
				Title("DNS mode").
				Options(
					huh.NewOption("Proxied (Cloudflare edge)", DNSProxied),
					huh.NewOption("DNS only", DNSOnly),
				).
				Value(&s.DNS),
		))
	}
	if s.Network == "" {
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[NetworkMode]().
				Title("Network mode").
				Options(
					huh.NewOption("IPv4", NetIPv4),
					huh.NewOption("IPv6", NetIPv6),
					huh.NewOption("Dual stack", NetDualStack),
				).
				Value(&s.Network),
		))
	}
	groups = append(groups, huh.NewGroup(
		huh.NewConfirm().
			Title("Allow NAT Gateways?").
			Description("Off by default — NAT is disallowed unless you flip this on.").
			Value(&s.AllowNAT),
	))

	if len(groups) == 0 {
		return nil
	}
	if err := huh.NewForm(groups...).Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}
	return nil
}
