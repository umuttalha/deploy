package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// Prompt fills any zero-valued fields on s by showing every option on a
// single huh form page. Navigate with Tab / Shift+Tab or arrow keys; submit
// to finish. Fields already populated (e.g. by cobra flags) are not shown.
func Prompt(s *Stack) error {
	fields := []huh.Field{
		huh.NewNote().
			Title("Configure stack").
			Description("Tab/↓ next · Shift+Tab/↑ back · Enter to confirm a select · Submit when done"),
	}

	if s.Name == "" {
		fields = append(fields,
			huh.NewInput().Title("Stack name").Value(&s.Name),
		)
	}
	if s.Region == "" {
		fields = append(fields,
			huh.NewInput().Title("AWS region").Placeholder("us-east-1").Value(&s.Region),
		)
	}
	if s.Storage == "" {
		fields = append(fields,
			huh.NewSelect[StorageType]().
				Title("Storage backend").
				Options(
					huh.NewOption("Amazon S3", StorageS3),
					huh.NewOption("Cloudflare R2", StorageR2),
				).
				Value(&s.Storage),
		)
	}
	if s.InstanceType == "" {
		fields = append(fields,
			huh.NewInput().Title("EC2 instance type").Placeholder("t3.micro").Value(&s.InstanceType),
		)
	}
	if s.Image.Registry == "" {
		fields = append(fields,
			huh.NewSelect[ImageRegistry]().
				Title("Image registry").
				Options(
					huh.NewOption("GitHub Container Registry", RegistryGHCR),
					huh.NewOption("Amazon ECR", RegistryECR),
				).
				Value(&s.Image.Registry),
		)
	}
	if s.Image.Repo == "" {
		fields = append(fields,
			huh.NewInput().Title("Image repo").Placeholder("owner/app").Value(&s.Image.Repo),
		)
	}
	if s.Image.Tag == "" {
		fields = append(fields,
			huh.NewInput().Title("Image tag").Placeholder("latest").Value(&s.Image.Tag),
		)
	}
	if s.DNS == "" {
		fields = append(fields,
			huh.NewSelect[DNSMode]().
				Title("DNS mode").
				Options(
					huh.NewOption("Proxied (Cloudflare edge)", DNSProxied),
					huh.NewOption("DNS only", DNSOnly),
				).
				Value(&s.DNS),
		)
	}
	if s.Network == "" {
		fields = append(fields,
			huh.NewSelect[NetworkMode]().
				Title("Network mode").
				Options(
					huh.NewOption("IPv4", NetIPv4),
					huh.NewOption("IPv6", NetIPv6),
					huh.NewOption("Dual stack", NetDualStack),
				).
				Value(&s.Network),
		)
	}
	fields = append(fields,
		huh.NewConfirm().
			Title("Allow NAT Gateways?").
			Description("Off by default — NAT is disallowed unless you flip this on.").
			Value(&s.AllowNAT),
	)

	if len(fields) <= 1 {
		// only the header note remains; nothing to ask.
		return nil
	}

	if err := huh.NewForm(huh.NewGroup(fields...)).Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}
	return nil
}
