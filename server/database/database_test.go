package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/malusev998/dusanmalusev/database"
	"github.com/malusev998/dusanmalusev/tests"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

type Cfg struct {
	conn string
}

func (Cfg) LazyConnect() bool {
	return true
}

func (Cfg) MaxConnLifetime() time.Duration {
	return 2 * time.Millisecond
}

func (Cfg) MaxConnIdleTime() time.Duration {
	return 2 * time.Millisecond
}

func (Cfg) HealthCheckPeriod() time.Duration {
	return 2 * time.Millisecond
}

func (Cfg) MaxConnections() int32 {
	return 5
}

func (Cfg) MinConnections() int32 {
	return 2
}

func (c Cfg) String() string {
	if c.conn == "" {
		return tests.GetConnectionString()
	}

	return c.conn
}

func TestConnectToDB_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool, err := database.ConnectDB(ctx, Cfg{}, log.Logger)

	assert.NoError(err)
	assert.NotNil(pool)
}

func TestConnectToDB_ConnectionStringError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool, err := database.ConnectDB(ctx, Cfg{conn: "Error string"}, log.Logger)

	assert.Error(err)
	assert.Nil(pool)
}

func TestConnectToDB_ConnectWithConfig(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := database.Config{
		URL:      tests.GetConnectionString(),
		MinConns: 2,
		MaxConns: 5,
		MaxConnectionLifetime: 2 * time.Millisecond,
		MaxConnectionIdleTime: 2 * time.Millisecond,
		HealthCheck: 2 * time.Millisecond,
	}

	pool, err := database.ConnectDB(ctx, cfg, log.Logger)

	assert.NoError(err)
	assert.NotNil(pool)
}

func TestConfig(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("URL", func(t *testing.T) {
		cfg := database.Config{
			URL: "url",
		}

		assert.Equal("url", cfg.String())
	})

	t.Run("Config_Generation", func(t *testing.T) {
		cfg := database.Config{
			Host:     "localhost",
			User:     "postgres",
			Password: "postgres",
			Port:     5432,
			DbName:   "test",
			TimeZone: "UTC",
			SslMode:  true,
		}

		assert.Equal("user=postgres host=localhost password=postgres dbname=test port=5432 sslmode=enable TimeZone=UTC", cfg.String())
	})
}
