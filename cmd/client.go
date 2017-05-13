package cmd

import (
	"github.com/felipejfc/go-tcp-example/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "init a client",
	Long:  `init a client`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &client.TCPClient{
			Host: host,
			Port: port,
		}
		client.Start()
	},
}

func init() {
	clientCmd.Flags().StringVarP(&host, "address", "a", "localhost", "the address of the server to connect to")
	clientCmd.Flags().IntVarP(&port, "port", "p", 10000, "the port of the server to connect to")
	RootCmd.AddCommand(clientCmd)
}
