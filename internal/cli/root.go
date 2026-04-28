package cli

import (
	"github.com/spf13/cobra"
)

type rootFlags struct {
	profile         string
	awsRegion       string
	cloudflareToken string
	nonInteractive  bool
}

var globalFlags rootFlags

var rootCmd = &cobra.Command{
	Use:           "deploy",
	Short:         "Provision a tagged AWS + Cloudflare stack",
	Long:          "deploy is a CLI scaffold for an infra provisioner. State is tracked via resource tags (managed-by=deploy, stack=<name>).",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&globalFlags.profile, "profile", "", "AWS profile name")
	rootCmd.PersistentFlags().StringVar(&globalFlags.awsRegion, "aws-region", "", "Default AWS region")
	rootCmd.PersistentFlags().StringVar(&globalFlags.cloudflareToken, "cloudflare-token", "", "Cloudflare API token (env: CLOUDFLARE_API_TOKEN)")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.nonInteractive, "non-interactive", false, "Fail if a value is missing instead of prompting")
}

// Execute is the binary entrypoint.
func Execute() error {
	return rootCmd.Execute()
}
