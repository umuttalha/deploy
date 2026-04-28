package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy <stack-name>",
	Short: "Tear down everything tagged stack=<name> (stub)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stackName := args[0]
		fmt.Fprintf(cmd.OutOrStdout(), "would destroy stack %q\n", stackName)
		fmt.Fprintln(cmd.OutOrStdout(), "(boilerplate — no resources were deleted)")
		// TODO: build awsx.Clients, cloudflarex.Clients, ghcr.Clients;
		// call DestroyEC2 / DestroyVPC / Destroy / DestroyECR / DestroyR2 / DestroyDNS / DestroyPackage
		// in reverse-dependency order.
		return nil
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
