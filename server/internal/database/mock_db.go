package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DBName = "postgres"
	DBUser = "postgres"
	DBPass = "postgres"
)

// TestContaierDatabase represents a test container database.
type TestContaierDatabase struct {
	DB        DB
	container testcontainers.Container
}

func SetupTestDatabase() (*TestContaierDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	container, dbInstance, err := createTestContainerDatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &TestContaierDatabase{
		container: container,
		DB:        dbInstance,
	}, nil
}

func (tdb *TestContaierDatabase) TearDown() {
	if tdb.DB != nil {
		tdb.DB.Close()
	}
	if err := testcontainers.TerminateContainer(tdb.container); err != nil {
		fmt.Println(err, "failed to terminate container")
	}
	// remove test container
	_ = tdb.container.Terminate(context.Background())

}

// Setup a postgres:16-alpine testdatabase using test container
// This will be used to test real db call.
func createTestContainerDatabase(ctx context.Context) (testcontainers.Container, DB, error) {
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(DBName),
		postgres.WithUsername(DBPass),
		postgres.WithPassword(DBUser),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("5432/tcp").
				WithStartupTimeout(60*time.Second),
		),
	)

	if err != nil {
		return nil, nil, err
	}
	connectionString, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		return nil, nil, err
	}

	db, err := NewDatabase(&DatabaseConfig{
		ConenctionString:        connectionString,
		DBMaxOpenConnections:    10,
		DBMaxConnectionLifeTime: "1m",
		DBMaxConnectionIdleTime: "30m",
	})
	if err != nil {
		return pgContainer, db, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return pgContainer, db, nil
}
