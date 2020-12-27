package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/malusev998/dusanmalusev/config"
)

func TestConfigWithIOReader(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	reader := strings.NewReader(`
debug: true # Debug enables PProf
locale: en

postgres:
  uri: ''
  host: localhost
  user: gofiber-boilerplate
  password: gofiber-boilerplate
  dbname: gofiber-boilerplate
  port: 5432
  timezone: UTC
  sslmode: false
  logfile: ./log/db.log

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

	assert.Empty(cfg.Database.URI)
	assert.Equal("localhost", cfg.Database.Host)
	assert.Equal("gofiber-boilerplate", cfg.Database.User)
	assert.Equal("gofiber-boilerplate", cfg.Database.Password)
	assert.Equal("gofiber-boilerplate", cfg.Database.DBName)
	assert.Equal(int32(5432), cfg.Database.Port)
	assert.Equal("UTC", cfg.Database.TimeZone)
	assert.Equal(false, cfg.Database.SSLMode)
	assert.Equal("./log/db.log", cfg.Database.LogFile)

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

	assert.Empty(cfg.Database.URI)
	assert.Equal("localhost", cfg.Database.Host)
	assert.Equal("gofiber-boilerplate", cfg.Database.User)
	assert.Equal("gofiber-boilerplate", cfg.Database.Password)
	assert.Equal("gofiber-boilerplate", cfg.Database.DBName)
	assert.Equal(int32(5432), cfg.Database.Port)
	assert.Equal("UTC", cfg.Database.TimeZone)
	assert.Equal(false, cfg.Database.SSLMode)
	assert.Equal("./log/db.log", cfg.Database.LogFile)

	assert.Equal("./log/server.log", cfg.Logging.File)
	assert.Equal(true, cfg.Logging.Console)
	assert.Equal("info", cfg.Logging.Level)

	assert.Equal(false, cfg.HTTP.Prefork)
	assert.Equal(":4000", cfg.HTTP.Address)
}
