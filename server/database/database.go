package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)


type ConfigInterface interface {
	fmt.Stringer
	LazyConnect() bool
	MaxConnLifetime() time.Duration
	MaxConnIdleTime() time.Duration
	HealthCheckPeriod() time.Duration
	MaxConnections() int32
	MinConnections() int32
}

type Config struct {
	URL                   string
	Host                  string
	User                  string
	Password              string
	DbName                string
	TimeZone              string
	MaxConnectionLifetime time.Duration
	MaxConnectionIdleTime time.Duration
	HealthCheck           time.Duration
	MaxConns              int32
	MinConns              int32
	Port                  uint16
	SslMode               bool
	Lazy                  bool
}

func (c Config) LazyConnect() bool {
	return c.Lazy
}

func (c Config) MaxConnLifetime() time.Duration {
	return c.MaxConnectionLifetime
}

func (c Config) MaxConnIdleTime() time.Duration {
	return c.MaxConnectionIdleTime
}

func (c Config) HealthCheckPeriod() time.Duration {
	return c.HealthCheck
}

func (c Config) MaxConnections() int32 {
	return c.MaxConns
}

func (c Config) MinConnections() int32 {
	return c.MinConns
}

func (c Config) String() string {
	if c.URL != "" {
		return c.URL
	}

	sslMode := "disable"

	if c.SslMode {
		sslMode = "enable"
	}

	return fmt.Sprintf(
		"user=%s host=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.User,
		c.Host,
		c.Password,
		c.DbName,
		c.Port,
		sslMode,
		c.TimeZone,
	)
}

func ConnectDB(ctx context.Context, c ConfigInterface, logger zerolog.Logger) (pool *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig(c.String())
	if err != nil {
		return nil, err
	}

	config.LazyConnect = c.LazyConnect()
	config.MaxConnLifetime = c.MaxConnLifetime()
	config.MaxConnIdleTime = c.MaxConnIdleTime()
	config.HealthCheckPeriod = c.HealthCheckPeriod()
	config.MaxConns = c.MaxConnections()
	config.MinConns = c.MinConnections()

	config.ConnConfig.Logger = zerologadapter.NewLogger(logger)
	pool, err = pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		return nil, err
	}

	return pool, err
}
