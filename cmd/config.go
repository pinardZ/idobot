package cmd

import (
	"github.com/spf13/viper"
	"log"
)

var _ Command = (*configCommand)(nil)

type configCommand struct{}

func newConfigCmd() *configCommand {
	return &configCommand{}
}

func (c *configCommand) Execute() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.SetConfigFile("conf/config.yml")
	return viper.ReadInConfig()
}
