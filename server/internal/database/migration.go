package database

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// MigrationFS represents migration filesystem.
//
//go:embed migrations/*.sql
var MigrationFS embed.FS

// MigrationDIR where db migrations exists
var MigrationDIR = "migrations"

// runDBMigrations runs db migrations which exists in server/internal/database/migrations
func runDBMigrations(cfg *DatabaseConfig) error {
	pgxConfig, err := pgxpool.ParseConfig(cfg.ConenctionString)
	if err != nil {
		return err
	}
	database := stdlib.OpenDB(*pgxConfig.ConnConfig)
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return err
	}

	d, err := iofs.New(MigrationFS, MigrationDIR)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)

	if err != nil {
		return err
	}

	// Run the migration to upgrade the schema.
	err = m.Up()
	if err == migrate.ErrNoChange {
		cfg.Logger.Info("No change in migrations")
		return nil
	}
	if err != nil && err != migrate.ErrNoChange {
		cfg.Logger.Error(err, "unable to apply migrations")
		return err
	}
	cfg.Logger.Info("migrations applied successfully")
	return nil
}
