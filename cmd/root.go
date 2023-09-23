package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd is create new root command instance.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	rootCmd.AddCommand(NewServeCmd())
	return rootCmd
}
