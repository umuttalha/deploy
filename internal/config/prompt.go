package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// Prompt fills any zero-valued fields on s by walking the user through a
// multi-section huh form.
//
// Navigation:
//   - Tab / Enter   → next field or section
//   - Shift+Tab     → previous field or section
//   - Esc / Ctrl+C  → cancel
//
// Fields already populated (e.g. by cobra flags) are skipped within their
// section. A section with no remaining unset fields is skipped entirely.
func Prompt(s *Stack) error {
	var groups []*huh.Group

	// --- Section 1: Identity ---
	identity := []huh.Field{
		huh.NewNote().Title("Section 1 of 5 · Identity"),
	}
	if s.Name == "" {
		identity = append(identity,
			huh.NewInput().Title("Stack name").Value(&s.Name),
		)
	}
	if s.Region == "" {
		identity = append(identity,
			huh.NewInput().Title("AWS region").Placeholder("us-east-1").Value(&s.Region),
		)
	}
	if len(identity) > 1 {
		groups = append(groups, huh.NewGroup(identity...))
	}

	// --- Section 2: Storage ---
	storage := []huh.Field{
		huh.NewNote().Title("Section 2 of 5 · Storage"),
	}
	if s.Storage == "" {
		storage = append(storage,
			huh.NewSelect[StorageType]().
				Title("Storage backend").
				Options(
					huh.NewOption("Amazon S3", StorageS3),
					huh.NewOption("Cloudflare R2", StorageR2),
				).
				Value(&s.Storage),
		)
	}
	if len(storage) > 1 {
		groups = append(groups, huh.NewGroup(storage...))
	}

	// --- Section 3: Compute ---
	compute := []huh.Field{
		huh.NewNote().Title("Section 3 of 5 · Compute"),
	}
	if s.InstanceType == "" {
		compute = append(compute,
			huh.NewInput().Title("EC2 instance type").Placeholder("t3.micro").Value(&s.InstanceType),
		)
	}
	compute = append(compute,
		huh.NewConfirm().
			Title("Allow NAT Gateways?").
			Description("Off by default — NAT is disallowed unless you flip this on.").
			Value(&s.AllowNAT),
	)
	if len(compute) > 1 {
		groups = append(groups, huh.NewGroup(compute...))
	}

	// --- Section 4: Container image ---
	image := []huh.Field{
		huh.NewNote().Title("Section 4 of 5 · Container image"),
	}
	if s.Image.Registry == "" {
		image = append(image,
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
		image = append(image,
			huh.NewInput().Title("Image repo").Placeholder("owner/app").Value(&s.Image.Repo),
		)
	}
	if s.Image.Tag == "" {
		image = append(image,
			huh.NewInput().Title("Image tag").Placeholder("latest").Value(&s.Image.Tag),
		)
	}
	if len(image) > 1 {
		groups = append(groups, huh.NewGroup(image...))
	}

	// --- Section 5: Networking ---
	networking := []huh.Field{
		huh.NewNote().Title("Section 5 of 5 · Networking"),
	}
	if s.DNS == "" {
		networking = append(networking,
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
		networking = append(networking,
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
	if len(networking) > 1 {
		groups = append(groups, huh.NewGroup(networking...))
	}

	if len(groups) == 0 {
		return nil
	}
	if err := huh.NewForm(groups...).Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}
	return nil
}
