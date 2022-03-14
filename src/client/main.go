package main

import (
	"os"

	"github.com/alecthomas/kong"

	"portService/client/cli"
)

var CLI struct {
	Poll cli.PollCmd `cmd:"" help:"Poll GRPC server on an interval"`
}

func main() {
	os.Args = []string{"prot-service-client", "poll"}
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&cli.Context{})
	ctx.FatalIfErrorf(err)

}
