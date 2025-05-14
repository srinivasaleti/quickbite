package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/config"
	"github.com/srinivasaleti/quickbite/server/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	t.Run("should return error on invalid command", func(t *testing.T) {
		rootCmd.SetArgs([]string{"invalid"})
		err := rootCmd.Execute()
		assert.Error(t, err)
	})

	t.Run("should start server", func(t *testing.T) {
		t.Setenv("QUICKBITE_CONNECTION_STRING", `postgres://postgres:postgres@server:15435/postgres?sslmode=disable`)
		var mockServer = server.MockServer{}
		mockServer.On("Start")
		createServerFunc = func(p string, configuration config.ServerConfiguration) (server.IServer, error) {
			assert.Equal(t, "1234", p)
			return &mockServer, nil
		}

		rootCmd.SetArgs([]string{"server", "--port=1234"})
		err := rootCmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, "1234", port)
		mockServer.AssertCalled(t, "Start")
	})

	t.Run("should not start server if unable to create", func(t *testing.T) {
		var mockServer = server.MockServer{}
		createServerFunc = func(p string, configuration config.ServerConfiguration) (server.IServer, error) {
			return nil, fmt.Errorf("server creation failed")
		}

		rootCmd.SetArgs([]string{"server", "--port=4321"})
		err := rootCmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, "4321", port)
		mockServer.AssertNotCalled(t, "Start")
	})

	t.Run("should not start server when config not loaded properly", func(t *testing.T) {
		os.Unsetenv("QUICKBITE_CONNECTION_STRING")
		var mockServer = server.MockServer{}
		createServerFunc = func(p string, configuration config.ServerConfiguration) (server.IServer, error) {
			return nil, fmt.Errorf("server creation failed")
		}

		rootCmd.SetArgs([]string{"server", "--port=4321"})
		err := rootCmd.Execute()

		assert.NoError(t, err)
		assert.Equal(t, "4321", port)
		mockServer.AssertNotCalled(t, "Start")
	})
}
