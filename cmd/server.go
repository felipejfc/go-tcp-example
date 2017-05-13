package cmd

import (
	"github.com/felipejfc/tcp-example/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "init a server",
	Long:  `init a server`,
	Run: func(cmd *cobra.Command, args []string) {
		server := &server.TCPServer{
			Bind: host,
			Port: port,
		}
		server.Start()

	},
}

func init() {
	serverCmd.Flags().StringVarP(&host, "bind", "b", "localhost", "the address to bind")
	serverCmd.Flags().IntVarP(&port, "port", "p", 10000, "the port to bind")
	RootCmd.AddCommand(serverCmd)
}
