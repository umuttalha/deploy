package config

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"

	"github.com/umuttalha/deploy/internal/ui"
)

// customMarker is the sentinel value selected from a preset list when the
// user wants to type their own value.
const customMarker = "__custom__"

// minimalTheme returns a huh theme stripped of all color — plain white text,
// no accents, no bold highlights. Mimics Vercel CLI's clean prompt style.
func minimalTheme() *huh.Theme {
	t := huh.ThemeBase()

	plain := lipgloss.NewStyle()

	t.Focused.Title = plain
	t.Focused.Description = plain.Faint(true)
	t.Focused.Base = plain
	t.Focused.SelectSelector = plain.SetString("> ")
	t.Focused.SelectedOption = plain
	t.Focused.UnselectedOption = plain.Faint(true)
	t.Focused.TextInput.Cursor = plain
	t.Focused.TextInput.Placeholder = plain.Faint(true)
	t.Focused.TextInput.Prompt = plain.SetString("> ")

	t.Blurred = t.Focused

	return t
}

// promptSelect runs a single huh select prompt and prints the answered line.
func promptSelect[T comparable](title string, opts []huh.Option[T], val *T) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[T]().
				Title(title).
				Options(opts...).
				Value(val),
		),
	).WithTheme(minimalTheme())

	if err := form.Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}
	ui.Answered(os.Stdout, title, fmt.Sprintf("%v", *val))
	return nil
}

// promptInput runs a single huh text input and prints the answered line.
func promptInput(title, placeholder string, val *string) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(title).
				Placeholder(placeholder).
				Value(val),
		),
	).WithTheme(minimalTheme())

	if err := form.Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}
	ui.Answered(os.Stdout, title, *val)
	return nil
}

// promptConfirm runs a single huh confirm and prints the answered line.
func promptConfirm(title string, val *bool) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Value(val),
		),
	).WithTheme(minimalTheme())

	if err := form.Run(); err != nil {
		return fmt.Errorf("prompt: %w", err)
	}
	label := "no"
	if *val {
		label = "yes"
	}
	ui.Answered(os.Stdout, title, label)
	return nil
}

// Prompt fills any zero-valued fields on s by showing one prompt at a time
// in the style of Vercel CLI. Fields already set (e.g. via cobra flags) are
// skipped entirely.
func Prompt(s *Stack) error {
	// Stack name
	if s.Name == "" {
		if err := promptInput("What's your stack name?", "my-stack", &s.Name); err != nil {
			return err
		}
	}

	// AWS region
	if s.Region == "" {
		var choice string
		if err := promptSelect("AWS region", []huh.Option[string]{
			huh.NewOption("us-east-1 (N. Virginia)", "us-east-1"),
			huh.NewOption("us-west-2 (Oregon)", "us-west-2"),
			huh.NewOption("eu-west-1 (Ireland)", "eu-west-1"),
			huh.NewOption("Custom…", customMarker),
		}, &choice); err != nil {
			return err
		}
		if choice == customMarker {
			if err := promptInput("Custom AWS region", "e.g. ap-southeast-2", &s.Region); err != nil {
				return err
			}
		} else {
			s.Region = choice
		}
	}

	// Storage backend
	if s.Storage == "" {
		if err := promptSelect("Storage backend", []huh.Option[StorageType]{
			huh.NewOption("Amazon S3", StorageS3),
			huh.NewOption("Cloudflare R2", StorageR2),
		}, &s.Storage); err != nil {
			return err
		}
	}

	// EC2 instance type
	if s.InstanceType == "" {
		var choice string
		if err := promptSelect("EC2 instance type", []huh.Option[string]{
			huh.NewOption("t3.micro (2 vCPU · 1 GiB · burstable)", "t3.micro"),
			huh.NewOption("t3.small (2 vCPU · 2 GiB · burstable)", "t3.small"),
			huh.NewOption("t3.medium (2 vCPU · 4 GiB · burstable)", "t3.medium"),
			huh.NewOption("Custom…", customMarker),
		}, &choice); err != nil {
			return err
		}
		if choice == customMarker {
			if err := promptInput("Custom EC2 instance type", "e.g. m5.large", &s.InstanceType); err != nil {
				return err
			}
		} else {
			s.InstanceType = choice
		}
	}

	// Image registry
	if s.Image.Registry == "" {
		if err := promptSelect("Image registry", []huh.Option[ImageRegistry]{
			huh.NewOption("GitHub Container Registry", RegistryGHCR),
			huh.NewOption("Amazon ECR", RegistryECR),
		}, &s.Image.Registry); err != nil {
			return err
		}
	}

	// Image repo
	if s.Image.Repo == "" {
		if err := promptInput("Image repo", "owner/app", &s.Image.Repo); err != nil {
			return err
		}
	}

	// Image tag
	if s.Image.Tag == "" {
		if err := promptInput("Image tag", "latest", &s.Image.Tag); err != nil {
			return err
		}
	}

	// DNS mode
	if s.DNS == "" {
		if err := promptSelect("DNS mode", []huh.Option[DNSMode]{
			huh.NewOption("Proxied (Cloudflare edge)", DNSProxied),
			huh.NewOption("DNS only", DNSOnly),
		}, &s.DNS); err != nil {
			return err
		}
	}

	// Network mode
	if s.Network == "" {
		if err := promptSelect("Network mode", []huh.Option[NetworkMode]{
			huh.NewOption("IPv4", NetIPv4),
			huh.NewOption("IPv6", NetIPv6),
			huh.NewOption("Dual stack", NetDualStack),
		}, &s.Network); err != nil {
			return err
		}
	}

	// NAT
	if err := promptConfirm("Allow NAT Gateways?", &s.AllowNAT); err != nil {
		return err
	}

	return nil
}

func requireNonEmpty(label string) func(string) error {
	return func(v string) error {
		if v == "" {
			return fmt.Errorf("%s is required", label)
		}
		return nil
	}
}
