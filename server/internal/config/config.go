package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/srinivasaleti/quickbite/server/internal/database"
)

// DBConfiguration defines the db configuration.
type DBConfiguration struct {
	// ConnectionString represents database connection string.
	ConnectionString string `required:"true" envconfig:"QUICKBITE_CONNECTION_STRING"`
	// MaxConnections is the maximum size of the pool.
	MaxConnections int `required:"false" envconfig:"QUICKBITE_DB_MAX_CONNECTIONS" default:"30"`
	// ConnectionLifeTime is the duration since creation after which a connection will be automatically closed.
	ConnectionLifeTime string `required:"false" envconfig:"QUICKBITE_DB_MAX_CONNECTIONS_LIFE_TIME" default:"30m"`
	// MaxConnectionIdleTime is the duration after which an idle connection will be automatically closed by the health check.
	MaxConnectionIdleTime string `required:"false" envconfig:"QUICKBITE_DB_MAX_CONNECTIONS_IDLE_TIME" default:"1m"`
	// DBSSLMode is represents postgres ssl mode.
	DBSSLMode string `required:"false" envconfig:"QUICKBITE_DB_SSL_MODE" default:"require"`
}

// ServerConfiguration defines the server configuration.
type ServerConfiguration struct {
	DBConfiguration
}

func (s *ServerConfiguration) DBConfig() *database.DatabaseConfig {
	return &database.DatabaseConfig{
		DBMaxOpenConnections:    s.MaxConnections,
		DBMaxConnectionLifeTime: s.ConnectionLifeTime,
		DBMaxConnectionIdleTime: s.MaxConnectionIdleTime,
		ConenctionString:        s.ConnectionString,
	}
}

// NewServerConfiguration returns a Server with default parameters.
func NewServerConfiguration() (*ServerConfiguration, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load")
	}
	var s ServerConfiguration
	if err := envconfig.Process("", &s); err != nil {
		return nil, err
	}
	return &s, nil
}
