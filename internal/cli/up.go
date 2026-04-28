package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/umuttalha/deploy/internal/config"
	"github.com/umuttalha/deploy/internal/ui"
	"github.com/umuttalha/deploy/internal/version"
)

var upStack config.Stack

var upCmd = &cobra.Command{
	Use:     "up [stack-name]",
	Aliases: []string{"deploy"},
	Short:   "Provision a stack",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		ui.Header(out, "Deploy CLI", version.Version)

		if len(args) == 1 {
			upStack.Name = args[0]
		}

		if !globalFlags.nonInteractive {
			if err := config.Prompt(&upStack); err != nil {
				return err
			}
		}

		if err := config.Validate(upStack); err != nil {
			hint := ""
			if globalFlags.nonInteractive {
				hint = "drop --non-interactive to fill missing fields, or pass them as flags"
			}
			ui.Fail(out, fmt.Sprintf("invalid stack: %s", err), hint)
			return errSilent
		}

		ui.OK(out, "Validated stack %q", upStack.Name)
		ui.KV(out, "Region", upStack.Region)
		ui.KV(out, "Storage", string(upStack.Storage))
		ui.KV(out, "Instance", upStack.InstanceType)
		ui.KV(out, "Image", fmt.Sprintf("%s / %s:%s", upStack.Image.Registry, upStack.Image.Repo, upStack.Image.Tag))
		ui.KV(out, "DNS", string(upStack.DNS))
		ui.KV(out, "Network", string(upStack.Network))
		ui.KV(out, "NAT", natLabel(upStack.AllowNAT))
		fmt.Fprintln(out)
		ui.Warn(out, "Boilerplate — no resources created")
		return nil
	},
}

func natLabel(allow bool) string {
	if allow {
		return "allowed"
	}
	return "disallowed"
}

func init() {
	f := upCmd.Flags()
	f.StringVar(&upStack.Region, "region", "", "AWS region")
	f.StringVar((*string)(&upStack.Storage), "storage", "", "Storage backend: s3 | r2")
	f.StringVar(&upStack.InstanceType, "instance-type", "", "EC2 instance type, e.g. t3.micro")
	f.StringVar((*string)(&upStack.Image.Registry), "image-registry", "", "Image registry: ghcr | ecr")
	f.StringVar(&upStack.Image.Repo, "image-repo", "", "Image repository, e.g. owner/app")
	f.StringVar(&upStack.Image.Tag, "image-tag", "", "Image tag, e.g. latest")
	f.StringVar((*string)(&upStack.DNS), "dns", "", "DNS mode: proxied | dns-only")
	f.StringVar((*string)(&upStack.Network), "network", "", "Network mode: ipv4 | ipv6 | dual")
	f.BoolVar(&upStack.AllowNAT, "allow-nat", false, "Allow NAT Gateways (off by default)")

	rootCmd.AddCommand(upCmd)
}
