package main

import (
	"github.com/alecthomas/kong"

	"portService/client/cli"
)

var CLI struct {
	Poll cli.PollCmd `cmd:"" help:"Poll GRPC server on an interval"`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&cli.Context{})
	ctx.FatalIfErrorf(err)

}
