package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/umuttalha/deploy/internal/config"
)

var deployStack config.Stack

var deployCmd = &cobra.Command{
	Use:   "deploy [stack-name]",
	Short: "Provision a stack (boilerplate — no resources are actually created yet)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			deployStack.Name = args[0]
		}

		if !globalFlags.nonInteractive {
			if err := config.Prompt(&deployStack); err != nil {
				return err
			}
		}

		if err := config.Validate(deployStack); err != nil {
			if globalFlags.nonInteractive {
				return fmt.Errorf("invalid stack: %w (run without --non-interactive to fill missing fields)", err)
			}
			return fmt.Errorf("invalid stack: %w", err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "resolved stack:\n%+v\n", deployStack)
		fmt.Fprintln(cmd.OutOrStdout(), "(boilerplate — no resources were created)")
		return nil
	},
}

func init() {
	f := deployCmd.Flags()
	f.StringVar(&deployStack.Region, "region", "", "AWS region")
	f.StringVar((*string)(&deployStack.Storage), "storage", "", "Storage backend: s3 | r2")
	f.StringVar(&deployStack.InstanceType, "instance-type", "", "EC2 instance type, e.g. t3.micro")
	f.StringVar((*string)(&deployStack.Image.Registry), "image-registry", "", "Image registry: ghcr | ecr")
	f.StringVar(&deployStack.Image.Repo, "image-repo", "", "Image repository, e.g. owner/app")
	f.StringVar(&deployStack.Image.Tag, "image-tag", "", "Image tag, e.g. latest")
	f.StringVar((*string)(&deployStack.DNS), "dns", "", "DNS mode: proxied | dns-only")
	f.StringVar((*string)(&deployStack.Network), "network", "", "Network mode: ipv4 | ipv6 | dual")
	f.BoolVar(&deployStack.AllowNAT, "allow-nat", false, "Allow NAT Gateways (off by default)")

	rootCmd.AddCommand(deployCmd)
}
