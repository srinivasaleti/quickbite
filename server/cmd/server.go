package cmd

import (
	"quickbite/server/internal/server"

	"github.com/spf13/cobra"
)

var port string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		server.NewServer(port)
	},
}

func init() {
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	rootCmd.AddCommand(serverCmd)
}
