package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

const DefaultDBOperationTimeout = time.Second * 5

type DB interface {
	Close()
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
	Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error)
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
