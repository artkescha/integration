package cmd

import (
	"github.com/spf13/cobra"

	"github.com/artkescha/integration/cmd/client"
	"github.com/artkescha/integration/cmd/server"
)

var Command = &cobra.Command{
	Use:   "integration",
	Short: "this is grpc demo with integration tests",
}

func init() {
	Command.AddCommand(server.Command)
	Command.AddCommand(client.Command)
}
