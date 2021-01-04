package config

import (
	"fmt"
	"io"
	"time"

	"github.com/spf13/viper"
)

type (
	DB struct {
		URI                   string        `json:"uri,omitempty" yaml:"uri,omitempty"`
		Host                  string        `json:"host,omitempty" yaml:"host,omitempty"`
		User                  string        `json:"user,omitempty" yaml:"user,omitempty"`
		Password              string        `json:"password,omitempty" yaml:"password,omitempty"`
		DBName                string        `json:"dbname,omitempty" yaml:"dbname,omitempty"`
		TimeZone              string        `json:"timezone,omitempty" yaml:"timezone,omitempty"`
		LogFile               string        `json:"logfile,omitempty" yaml:"logfile,omitempty"`
		MaxConns              int32         `json:"max_conns,omitempty" yaml:"max_conns,omitempty"`
		MinConns              int32         `json:"min_conns,omitempty" yaml:"min_conns,omitempty"`
		Port                  int32         `json:"port,omitempty" yaml:"port,omitempty"`
		SSLMode               bool          `json:"sslmode,omitempty" yaml:"sslmode,omitempty"`
		Lazy                  bool          `json:"lazy,omitempty" yaml:"lazy,omitempty"`
		HealthCheck           time.Duration `json:"health_check,omitempty" yaml:"health_check,omitempty"`
		MaxConnectionIdleTime time.Duration `json:"max_connection_idle_time,omitempty" yaml:"max_connection_idle_time,omitempty"`
		MaxConnectionLifetime time.Duration `json:"max_connection_lifetime,omitempty" yaml:"max_connection_lifetime,omitempty"`
	}

	Logging struct {
		Level   string `json:"level" yaml:"level"`
		File    string `json:"file" yaml:"file"`
		Console bool   `json:"console" yaml:"console"`
	}

	HTTP struct {
		Address string `json:"address" yaml:"address"`
		Prefork bool   `json:"prefork" yaml:"prefork"`
	}

	Subscription struct {
		SendEmail bool `json:"send_email,omitempty" yaml:"send_email,omitempty"`
	}

	SMTP struct {
		Address  string `json:"address,omitempty" yaml:"address,omitempty"`
		From     string `json:"from,omitempty" yaml:"from,omitempty"`
		PoolSize int    `json:"pool_size,omitempty" yaml:"pool_size,omitempty"`
		Senders  int    `json:"senders,omitempty" yaml:"senders,omitempty"`
	}

	Config struct {
		Database DB      `json:"postgres" yaml:"postgres"`
		Logging  Logging `json:"logging" yaml:"logging"`
		HTTP     HTTP    `json:"http" yaml:"http"`

		Locale       string       `json:"locale" yaml:"locale"`
		Subscription Subscription `json:"subscription,omitempty" yaml:"subscription,omitempty"`
		Debug        bool         `json:"debug" yaml:"debug"`
	}
)

func New(name string, paths ...interface{}) (config Config, err error) {
	viperInstance := viper.New()
	viperInstance.SetConfigType("yaml")
	viperInstance.AutomaticEnv()

	for _, path := range paths {
		switch pathType := path.(type) {
		case string:
			viperInstance.SetConfigName(name)
			viperInstance.AddConfigPath(pathType)
			err = viperInstance.ReadInConfig()
			if err != nil {
				return Config{}, err
			}
		case io.Reader:
			if err := viperInstance.ReadConfig(pathType); err != nil {
				return Config{}, err
			}
		}
	}

	fmt.Println(viperInstance.GetString("postgres.uri"))

	err = viperInstance.Unmarshal(&config)

	if err != nil {
		return
	}

	config.Database = DB{
		URI:                   viperInstance.GetString("postgres.uri"),
		Host:                  viperInstance.GetString("postgres.host"),
		User:                  viperInstance.GetString("postgres.user"),
		Password:              viperInstance.GetString("postgres.password"),
		Port:                  viperInstance.GetInt32("postgres.port"),
		DBName:                viperInstance.GetString("postgres.dbname"),
		TimeZone:              viperInstance.GetString("postgres.timezone"),
		SSLMode:               viperInstance.GetBool("postgres.sslmode"),
		LogFile:               viperInstance.GetString("postgres.logfile"),
		MaxConnectionLifetime: viperInstance.GetDuration("postgres.max_connection_lifetime"),
		MaxConnectionIdleTime: viperInstance.GetDuration("postgres.max_connection_idle_time"),
		HealthCheck:           viperInstance.GetDuration("postgres.health_check"),
		MaxConns:              viperInstance.GetInt32("postgres.max_conns"),
		MinConns:              viperInstance.GetInt32("postgres.min_conns"),
		Lazy:                  viperInstance.GetBool("postgres.lazy"),
	}

	return
}
