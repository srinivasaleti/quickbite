package config

import (
	"os"
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestNewServerConfiguration(t *testing.T) {
	t.Run("should return error when required config not provided", func(t *testing.T) {
		os.Unsetenv("QUICKBITE_DB_CONNECTION_STRING")
		s, err := NewServerConfiguration()

		assert.Error(t, err)
		assert.Nil(t, s)
	})
	t.Run("should return config", func(t *testing.T) {
		t.Setenv("QUICKBITE_DB_CONNECTION_STRING", `postgres://postgres:postgres@server:15435/postgres?sslmode=disable`)
		s, err := NewServerConfiguration()

		assert.NoError(t, err)
		assert.Equal(t, s.ConnectionString, `postgres://postgres:postgres@server:15435/postgres?sslmode=disable`)
	})
}

func TestDBConfiguration(t *testing.T) {
	t.Setenv("QUICKBITE_DB_CONNECTION_STRING", `postgres://postgres:postgres@server:15435/postgres?sslmode=disable`)
	s, err := NewServerConfiguration()

	assert.NoError(t, err)
	assert.Equal(t, s.DBConfig(), &database.DatabaseConfig{
		ConenctionString:        `postgres://postgres:postgres@server:15435/postgres?sslmode=disable`,
		DBMaxOpenConnections:    30,
		DBMaxConnectionLifeTime: "30m",
		DBMaxConnectionIdleTime: "1m",
	})

}
