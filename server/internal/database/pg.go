package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

// CreatePostgresDBPool creates a connection pool for postgres database.
func (c *DatabaseConfig) CreatePostgresDBPool(ctx context.Context) (*PostgresDB, error) {
	pgxConfig, err := pgxpool.ParseConfig(c.ConenctionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}
	dbMaxLifeTimeDuration, _ := time.ParseDuration(c.DBMaxConnectionLifeTime)
	dbMaxConnIdleTimeDuration, _ := time.ParseDuration(c.DBMaxConnectionIdleTime)
	pgxConfig.MaxConns = int32(c.DBMaxOpenConnections)
	pgxConfig.MaxConnLifetime = dbMaxLifeTimeDuration
	pgxConfig.MaxConnIdleTime = dbMaxConnIdleTimeDuration
	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		c.Logger.Error(err, "unable to create pgxpool")
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return &PostgresDB{Pool: pool}, nil
}

// PostgresDB should implment DB methods
func (db *PostgresDB) Close() {
	db.Pool.Close()
}

func (db *PostgresDB) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return db.Pool.SendBatch(ctx, b)
}

func (db *PostgresDB) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	return db.Pool.QueryRow(ctx, sql, arguments...)
}

func (db *PostgresDB) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	return db.Pool.Query(ctx, sql, arguments...)
}

func (db *PostgresDB) Begin(ctx context.Context) (pgx.Tx, error) {
	return db.Pool.Begin(ctx)
}

func ErrIsConstraint(err error, constraintName string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.ConstraintName == constraintName
	}
	return false
}
