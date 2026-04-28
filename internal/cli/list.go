package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List every stack tagged managed-by=deploy (stub)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(cmd.OutOrStdout(), "would list stacks tagged managed-by=deploy")
		fmt.Fprintln(cmd.OutOrStdout(), "(boilerplate — no resources were queried)")
		// TODO: awsx.New + cloudflarex.New, then call FindEC2/FindVPC/Find/FindECR/FindR2/FindDNS/FindPackage
		// across all stacks (no stack filter), group results by the "stack" tag, and print as a table.
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
