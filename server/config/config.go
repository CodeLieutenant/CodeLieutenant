package config

import (
	"encoding/base64"
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

	Redis struct {
		Host     string `json:"host,omitempty" yaml:"host,omitempty"`
		Username string `json:"username,omitempty" yaml:"username,omitempty"`
		Password string `json:"password,omitempty" yaml:"password,omitempty"`
		Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	}

	Csrf struct {
		CookieDomain string `json:"cookie_domain,omitempty" yaml:"cookie_domain,omitempty"`
		Secure       bool   `json:"secure,omitempty" yaml:"secure,omitempty"`
	}

	Session struct {
		CookieName   string        `json:"cookie_name,omitempty" yaml:"cookie_name,omitempty"`
		CookieDomain string        `json:"cookie_domain,omitempty" yaml:"cookie_domain,omitempty"`
		CookiePath   string        `json:"cookie_path,omitempty" yaml:"cookie_path,omitempty"`
		Secure       bool          `json:"secure,omitempty" yaml:"secure,omitempty"`
		Expiration   time.Duration `json:"expiration,omitempty" yaml:"expiration,omitempty"`
	}

	HTTP struct {
		Address string        `json:"address" yaml:"address"`
		Prefork bool          `json:"prefork" yaml:"prefork"`
		Timeout time.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
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
		Redis    Redis   `json:"redis" yaml:"redis"`
		Session  Session `json:"session" yaml:"session"`
		Csrf     Csrf    `json:"csrf" yaml:"csrf"`

		Locale       string       `json:"locale" yaml:"locale"`
		Subscription Subscription `json:"subscription,omitempty" yaml:"subscription,omitempty"`
		Debug        bool         `json:"debug" yaml:"debug"`
		Key          []byte       `json:"-" yaml:"-" mapstructure:"-"`
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

	config.HTTP = HTTP{
		Address: viperInstance.GetString("http.address"),
		Prefork: viperInstance.GetBool("http.prefork"),
		Timeout: viperInstance.GetDuration("http.timeout"),
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
		MaxConnectionLifetime: viperInstance.GetDuration("postgres.max_connection_lifetime"),
		MaxConnectionIdleTime: viperInstance.GetDuration("postgres.max_connection_idle_time"),
		HealthCheck:           viperInstance.GetDuration("postgres.health_check"),
		MaxConns:              viperInstance.GetInt32("postgres.max_conns"),
		MinConns:              viperInstance.GetInt32("postgres.min_conns"),
		Lazy:                  viperInstance.GetBool("postgres.lazy"),
	}

	config.Session = Session{
		CookieName:   viperInstance.GetString("session.cookie_name"),
		CookieDomain: viperInstance.GetString("session.cookie_domain"),
		CookiePath:   viperInstance.GetString("session.cookie_path"),
		Secure:       viperInstance.GetBool("session.secure"),
		Expiration:   viperInstance.GetDuration("session.expiration"),
	}

	config.Csrf = Csrf{
		CookieDomain: viperInstance.GetString("csrf.cookie_domain"),
		Secure:       viperInstance.GetBool("csrf.secure"),
	}

	config.Key, err = base64.RawURLEncoding.DecodeString(viperInstance.GetString("key"))

	return
}
