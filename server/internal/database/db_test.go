package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateDBPoolFromEnv(t *testing.T) {
	c := DatabaseConfig{
		ConenctionString:        "postgres://postgres:postgres@localhost:15431/postgres?sslmode=disable",
		DBMaxOpenConnections:    30,
		DBMaxConnectionLifeTime: "30m",
		DBMaxConnectionIdleTime: "1m",
	}
	dbPool, err := c.CreatePostgresDBPool(context.Background())
	dbMaxLifeTimeDuration, _ := time.ParseDuration("30m")
	dbMaxIdleLifeTimeDuration, _ := time.ParseDuration("1m")

	poolConfig := dbPool.Pool.Config()
	assert.Equal(t, poolConfig.MaxConns, int32(30))
	assert.Equal(t, poolConfig.MaxConnLifetime, dbMaxLifeTimeDuration)
	assert.Equal(t, poolConfig.MaxConnIdleTime, dbMaxIdleLifeTimeDuration)
	assert.NoError(t, err)

	dbPool.Close()
}

func TestNewDatabase(t *testing.T) {
	db, err := SetupTestDatabase()
	assert.NoError(t, err)
	assert.NotNil(t, db)
	db.TearDown()
}
