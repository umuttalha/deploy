package cli

import (
	"github.com/spf13/cobra"

	"github.com/umuttalha/deploy/internal/ui"
	"github.com/umuttalha/deploy/internal/version"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List every stack tagged managed-by=deploy",
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		ui.Header(out, "Deploy CLI", version.Version)
		ui.Info(out, "Querying resources tagged managed-by=deploy")
		ui.Info(out, "Boilerplate — no resources were queried")
		// TODO: awsx.New + cloudflarex.New, then call FindEC2/FindVPC/Find/FindECR/FindR2/FindDNS/FindPackage
		// across all stacks (no stack filter), group results by the "stack" tag, and print as a table.
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
