package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
)

type Command interface {
	Execute() error
}

var _ Command = (*command)(nil)

type command struct {
	ctx       context.Context
	configCmd *configCommand
	rootCmd   *cobra.Command
}

func (cmd command) Execute() error {
	err := cmd.configCmd.Execute()
	if err != nil {
		log.Fatal("init config before sniper")
	}
	return cmd.rootCmd.ExecuteContext(cmd.ctx)
}

func NewCommand(ctx context.Context) Command {
	cmd := &command{
		ctx:       ctx,
		configCmd: newConfigCmd(),
		rootCmd:   newRootCmd(),
	}
	return cmd
}
