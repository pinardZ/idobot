package cmd

import (
	"fmt"
	"github.com/pinardZ/idobot/eth"
	"github.com/spf13/cobra"
)

func newPreSaleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "presale",
		Short: "sniper presale",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello world")
			eth.NewDxSaleRunnable().Run(cmd.Context())
		},
	}
}
