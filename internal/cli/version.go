package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/umuttalha/deploy/internal/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(cmd.OutOrStdout(), version.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
