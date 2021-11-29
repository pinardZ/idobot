package main

import (
	"context"
	"github.com/pinardZ/idobot/cmd"
)

func main() {
	ctx := context.Background()
	cmd.NewCommand(ctx).Execute()
}
