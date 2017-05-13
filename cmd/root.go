package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var host string
var port int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-tcp-example",
	Short: "simple tcp server and client written in golang",
	Long:  `simple tcp server and client written in golang`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
}
