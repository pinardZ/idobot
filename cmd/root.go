package cmd

import "github.com/spf13/cobra"

const IDOBot = "ido-boot"

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   IDOBot,
		Short: IDOBot,
	}
	initRootCmd(rootCmd)
	return rootCmd
}

func initRootCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newPreSaleCmd())
}
