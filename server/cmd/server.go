package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srinivasaleti/planner/server/internal/config"
	"github.com/srinivasaleti/planner/server/internal/server"
)

var port string

var createServerFunc = server.NewServer

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		configuration, err := config.NewServerConfiguration()
		if err != nil {
			fmt.Println("unable to load configuration", err)
			return
		}
		s, err := createServerFunc(port, *configuration)
		if err != nil {
			fmt.Println("unable to start server", err)
			return
		}
		s.Start()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	rootCmd.AddCommand(serverCmd)
}
