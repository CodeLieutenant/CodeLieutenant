package config_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/config"
)

func TestConfigWithIOReader(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	reader := strings.NewReader(`
debug: true # Debug enables PProf
locale: en

postgres:
  host: localhost
  user: gofiber-boilerplate
  password: gofiber-boilerplate
  dbname: gofiber-boilerplate
  port: 5432
  timezone: UTC
  sslmode: false
  logfile: ./log/db.log
  max_connection_lifetime: 1h
  max_connection_idle_time: 24h
  health_check: 15m
  max_conns: 20
  min_conns: 5
  lazy: true

# supported logging - debug, info, warning, error
logging:
  level: info # Logger instance used in API
  console: true # Console logging
  file: ./log/server.log

http:
  prefork: false
  address: :4000 # HTTP Address
`)
	cfg, err := config.New("config", reader)

	assert.Nil(err)
	assert.NotZero(cfg)

	assert.True(cfg.Debug)
	assert.Equal("en", cfg.Locale)

	assert.Empty(cfg.Database.URI)
	assert.Equal("localhost", cfg.Database.Host)
	assert.Equal("gofiber-boilerplate", cfg.Database.User)
	assert.Equal("gofiber-boilerplate", cfg.Database.Password)
	assert.Equal("gofiber-boilerplate", cfg.Database.DBName)
	assert.Equal(int32(5432), cfg.Database.Port)
	assert.Equal("UTC", cfg.Database.TimeZone)
	assert.Equal(false, cfg.Database.SSLMode)
	assert.Equal(24*time.Hour, cfg.Database.MaxConnectionIdleTime)
	assert.Equal(1*time.Hour, cfg.Database.MaxConnectionLifetime)
	assert.Equal(15*time.Minute, cfg.Database.HealthCheck)
	assert.Equal(int32(20), cfg.Database.MaxConns)
	assert.Equal(int32(5), cfg.Database.MinConns)
	assert.Equal(true, cfg.Database.Lazy)

	assert.Equal("./log/server.log", cfg.Logging.File)
	assert.Equal(true, cfg.Logging.Console)
	assert.Equal("info", cfg.Logging.Level)

	assert.Equal(false, cfg.HTTP.Prefork)
	assert.Equal(":4000", cfg.HTTP.Address)
}

func TestConfigFile(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	cfg, err := config.New("config", "../tests")

	assert.Nil(err)
	assert.NotZero(cfg)

	assert.True(cfg.Debug)
	assert.Equal("en", cfg.Locale)

	assert.Empty(cfg.Database.URI)
	assert.Equal("localhost", cfg.Database.Host)
	assert.Equal("postgres", cfg.Database.User)
	assert.Equal("postgres", cfg.Database.Password)
	assert.Equal("dusanmalusev", cfg.Database.DBName)
	assert.Equal(int32(5432), cfg.Database.Port)
	assert.Equal("UTC", cfg.Database.TimeZone)
	assert.Equal(false, cfg.Database.SSLMode)
	assert.Equal(24*time.Hour, cfg.Database.MaxConnectionIdleTime)
	assert.Equal(1*time.Hour, cfg.Database.MaxConnectionLifetime)
	assert.Equal(15*time.Minute, cfg.Database.HealthCheck)
	assert.Equal(int32(20), cfg.Database.MaxConns)
	assert.Equal(int32(5), cfg.Database.MinConns)
	assert.Equal(true, cfg.Database.Lazy)

	assert.Equal("localhost", cfg.Redis.Host)
	assert.Equal(6379, cfg.Redis.Port)
	assert.Equal("", cfg.Redis.Password)
	assert.Equal("", cfg.Redis.Username)

	assert.Equal("dusanmalusev_session", cfg.Session.CookieName)
	assert.Equal("localhost", cfg.Session.CookieDomain)
	assert.Equal("", cfg.Session.CookiePath)
	assert.Equal(false, cfg.Session.Secure)
	assert.Equal(24*time.Hour, cfg.Session.Expiration)

	assert.Equal("localhost", cfg.Csrf.CookieDomain)
	assert.Equal(false, cfg.Csrf.Secure)

	assert.Equal("./log/server.log", cfg.Logging.File)
	assert.Equal(true, cfg.Logging.Console)
	assert.Equal("info", cfg.Logging.Level)

	assert.Equal(false, cfg.HTTP.Prefork)
	assert.Equal(":4000", cfg.HTTP.Address)
	assert.Equal(2*time.Second, cfg.HTTP.Timeout)
}
