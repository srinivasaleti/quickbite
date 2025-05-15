package database

import (
	"context"

	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type DB interface {
	Close()
}

type DatabaseConfig struct {
	Logger                  logger.ILogger
	ConenctionString        string
	DBMaxOpenConnections    int
	DBMaxConnectionLifeTime string
	DBMaxConnectionIdleTime string
}

type DBCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewDatabase(cfg *DatabaseConfig) (DB, error) {
	var err error
	db, err := cfg.CreatePostgresDBPool(context.Background())
	if err != nil {
		return nil, err
	}
	err = db.Pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	if err := runDBMigrations(cfg); err != nil {
		return nil, err
	}
	return db, err
}
