package cli

import (
	"github.com/spf13/cobra"

	"github.com/umuttalha/deploy/internal/ui"
	"github.com/umuttalha/deploy/internal/version"
)

var downCmd = &cobra.Command{
	Use:     "down <stack-name>",
	Aliases: []string{"destroy"},
	Short:   "Tear down a stack",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		stackName := args[0]

		ui.Header(out, "Deploy CLI", version.Version)
		ui.Info(out, "Tearing down stack %q", stackName)
		ui.Info(out, "Boilerplate — no resources deleted")
		// TODO: build awsx.Clients, cloudflarex.Clients, ghcr.Clients;
		// call DestroyEC2 / DestroyVPC / Destroy / DestroyECR / DestroyR2 / DestroyDNS / DestroyPackage
		// in reverse-dependency order.
		return nil
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
