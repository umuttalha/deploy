package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// customMarker is the sentinel value selected from a preset list when the
// user wants to type their own value.
const customMarker = "__custom__"

// Prompt fills any zero-valued fields on s by showing every option on a
// main huh form page. If the user picks "Custom…" for a preset field, a
// small follow-up page asks for the free-text value. Fields already set
// (e.g. via cobra flags) are skipped.
func Prompt(s *Stack) error {
	var (
		regionChoice   string
		regionCustom   string
		instanceChoice string
		instanceCustom string
	)

	main := []huh.Field{
		huh.NewNote().
			Title("Configure stack").
			Description("Tab/↓ next · Shift+Tab/↑ back · Enter on a select · Submit when done"),
	}

	if s.Name == "" {
		main = append(main,
			huh.NewInput().Title("Stack name").Value(&s.Name),
		)
	}
	if s.Region == "" {
		main = append(main,
			huh.NewSelect[string]().
				Title("AWS region").
				Options(
					huh.NewOption("us-east-1 (N. Virginia)", "us-east-1"),
					huh.NewOption("us-west-2 (Oregon)", "us-west-2"),
					huh.NewOption("eu-west-1 (Ireland)", "eu-west-1"),
					huh.NewOption("Custom…", customMarker),
				).
				Value(&regionChoice),
		)
	}
	if s.Storage == "" {
		main = append(main,
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
		main = append(main,
			huh.NewSelect[string]().
				Title("EC2 instance type").
				Options(
					huh.NewOption("t3.micro (2 vCPU · 1 GiB · burstable)", "t3.micro"),
					huh.NewOption("t3.small (2 vCPU · 2 GiB · burstable)", "t3.small"),
					huh.NewOption("t3.medium (2 vCPU · 4 GiB · burstable)", "t3.medium"),
					huh.NewOption("Custom…", customMarker),
				).
				Value(&instanceChoice),
		)
	}
	if s.Image.Registry == "" {
		main = append(main,
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
		main = append(main,
			huh.NewInput().Title("Image repo").Placeholder("owner/app").Value(&s.Image.Repo),
		)
	}
	if s.Image.Tag == "" {
		main = append(main,
			huh.NewInput().Title("Image tag").Placeholder("latest").Value(&s.Image.Tag),
		)
	}
	if s.DNS == "" {
		main = append(main,
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
		main = append(main,
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
	main = append(main,
		huh.NewConfirm().
			Title("Allow NAT Gateways?").
			Description("Off by default — NAT is disallowed unless you flip this on.").
			Value(&s.AllowNAT),
	)

	if len(main) <= 1 {
		// only the header note remains; nothing to ask.
		return nil
	}

	regionCustomGroup := huh.NewGroup(
		huh.NewInput().
			Title("Custom AWS region").
			Placeholder("e.g. ap-southeast-2").
			Validate(requireNonEmpty("region")).
			Value(&regionCustom),
	).WithHideFunc(func() bool { return regionChoice != customMarker })

	instanceCustomGroup := huh.NewGroup(
		huh.NewInput().
			Title("Custom EC2 instance type").
			Placeholder("e.g. m5.large").
			Validate(requireNonEmpty("instance type")).
			Value(&instanceCustom),
	).WithHideFunc(func() bool { return instanceChoice != customMarker })

	form := huh.NewForm(
		huh.NewGroup(main...),
		regionCustomGroup,
		instanceCustomGroup,
	)
	if err := form.Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}

	if s.Region == "" {
		s.Region = resolveCustom(regionChoice, regionCustom)
	}
	if s.InstanceType == "" {
		s.InstanceType = resolveCustom(instanceChoice, instanceCustom)
	}
	return nil
}

func resolveCustom(choice, custom string) string {
	if choice == customMarker {
		return custom
	}
	return choice
}

func requireNonEmpty(label string) func(string) error {
	return func(v string) error {
		if v == "" {
			return fmt.Errorf("%s is required", label)
		}
		return nil
	}
}
