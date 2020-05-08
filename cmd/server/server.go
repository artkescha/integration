package server

import (
	"fmt"

	"github.com/artkescha/integration/internal/server"
	"github.com/spf13/cobra"
)

var host string
var port int

func init() {
	Command.Flags().StringVar(&host, "host", "127.0.0.1", "host to listen")
	Command.Flags().IntVar(&port, "port", 8081, "port to listen")
}

var Command = &cobra.Command{
	Use:   "server",
	Short: "Run grpc server",
	Run: func(cmd *cobra.Command, args []string) {

		address := fmt.Sprintf("%s:%d", host, port)
		server.Run(cmd.Context(), address)
	},
}
