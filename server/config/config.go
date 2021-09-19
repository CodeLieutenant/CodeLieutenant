package config

import (
	"encoding/base64"
	"io"
	"time"

	"github.com/spf13/viper"
)

type (
	DB struct {
		URI                   string        `mapstructure:"uri" json:"uri,omitempty" yaml:"uri,omitempty"`
		Host                  string        `mapstructure:"host" json:"host,omitempty" yaml:"host,omitempty"`
		User                  string        `mapstructure:"user" json:"user,omitempty" yaml:"user,omitempty"`
		Password              string        `mapstructure:"password" json:"password,omitempty" yaml:"password,omitempty"`
		DBName                string        `mapstructure:"dbname" json:"dbname,omitempty" yaml:"dbname,omitempty"`
		TimeZone              string        `mapstructure:"timezone" json:"timezone,omitempty" yaml:"timezone,omitempty"`
		MaxConns              int32         `mapstructure:"max_conns" json:"max_conns,omitempty" yaml:"max_conns,omitempty"`
		MinConns              int32         `mapstructure:"min_conns" json:"min_conns,omitempty" yaml:"min_conns,omitempty"`
		Port                  int32         `mapstructure:"port" json:"port,omitempty" yaml:"port,omitempty"`
		SSLMode               bool          `mapstructure:"sslmode" json:"sslmode,omitempty" yaml:"sslmode,omitempty"`
		Lazy                  bool          `mapstructure:"lazy" json:"lazy,omitempty" yaml:"lazy,omitempty"`
		HealthCheck           time.Duration `mapstructure:"health_check" json:"health_check,omitempty" yaml:"health_check,omitempty"`
		MaxConnectionIdleTime time.Duration `mapstructure:"max_connection_idle_time" json:"max_connection_idle_time,omitempty" yaml:"max_connection_idle_time,omitempty"`
		MaxConnectionLifetime time.Duration `mapstructure:"max_connection_lifetime" json:"max_connection_lifetime,omitempty" yaml:"max_connection_lifetime,omitempty"`
	}

	Logging struct {
		Level   string `mapstructure:"level" json:"level,omitempty" yaml:"level,omitempty"`
		File    string `mapstructure:"file" json:"file,omitempty" yaml:"file,omitempty"`
		Console bool   `mapstructure:"console" json:"console,omitempty" yaml:"console,omitempty"`
	}

	Redis struct {
		Host     string `mapstructure:"host" json:"host,omitempty" yaml:"host,omitempty"`
		Username string `mapstructure:"username" json:"username,omitempty" yaml:"username,omitempty"`
		Password string `mapstructure:"password" json:"password,omitempty" yaml:"password,omitempty"`
		Port     int    `mapstructure:"port" json:"port,omitempty" yaml:"port,omitempty"`
	}

	Csrf struct {
		CookieDomain string `mapstructure:"cookie_domain" json:"cookie_domain,omitempty" yaml:"cookie_domain,omitempty"`
		Secure       bool   `mapstructure:"secure" json:"secure,omitempty" yaml:"secure,omitempty"`
	}

	Session struct {
		CookieName   string        `mapstructure:"cookie_name" json:"cookie_name,omitempty" yaml:"cookie_name,omitempty"`
		CookieDomain string        `mapstructure:"cookie_domain" json:"cookie_domain,omitempty" yaml:"cookie_domain,omitempty"`
		CookiePath   string        `mapstructure:"cookie_path" json:"cookie_path,omitempty" yaml:"cookie_path,omitempty"`
		Secure       bool          `mapstructure:"secure" json:"secure,omitempty" yaml:"secure,omitempty"`
		Expiration   time.Duration `mapstructure:"expiration" json:"expiration,omitempty" yaml:"expiration,omitempty"`
	}

	HTTP struct {
		Address string        `mapstructure:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
		Prefork bool          `mapstructure:"prefork,omitempty" json:"prefork,omitempty" yaml:"prefork,omitempty"`
		Timeout time.Duration `mapstructure:"timeout,omitempty" json:"timeout,omitempty" yaml:"timeout,omitempty"`
		BaseURL string        `mapstructure:"base_url,omitempty" json:"base_url,omitempty" yaml:"base_url,omitempty"`
	}

	Subscription struct {
		SendEmail bool `mapstructure:"send_email" json:"send_email,omitempty" yaml:"send_email,omitempty"`
	}

	SMTP struct {
		Host string `mapstructure:"address" json:"address,omitempty" yaml:"address,omitempty"`
		Port int    `mapstructure:"port" json:"port,omitempty" yaml:"port,omitempty"`
		From struct {
			Name  string `mapstructure:"name" json:"name,omitempty" yaml:"name,omitempty"`
			Email string `mapstructure:"email" json:"email,omitempty" yaml:"email,omitempty"`
		} `mapstructure:"from" json:"from,omitempty" yaml:"from,omitempty"`
		Username string `mapstructure:"username" json:"username,omitempty" yaml:"username,omitempty"`
		Password string `mapstructure:"password" json:"password,omitempty" yaml:"password,omitempty"`
		PoolSize int    `mapstructure:"pool_size" json:"pool_size,omitempty" yaml:"pool_size,omitempty"`
		Senders  int    `mapstructure:"senders" json:"senders,omitempty" yaml:"senders,omitempty"`
	}

	Config struct {
		Database DB      `json:"postgres" yaml:"postgres"`
		Logging  Logging `json:"logging" yaml:"logging"`
		HTTP     HTTP    `json:"http" yaml:"http"`
		Redis    Redis   `json:"redis" yaml:"redis"`
		Session  Session `json:"session" yaml:"session"`
		Csrf     Csrf    `json:"csrf" yaml:"csrf"`
		SMTP     SMTP    `json:"smtp" yaml:"smtp"`

		FrontendRedirectURL string       `json:"frontend_redirect_url,omitempty" yaml:"frontend_redirect_url,omitempty"`
		Locale              string       `json:"locale" yaml:"locale"`
		Subscription        Subscription `json:"subscription,omitempty" yaml:"subscription,omitempty"`
		Debug               bool         `json:"debug" yaml:"debug"`
		Key                 []byte       `json:"-" yaml:"-" mapstructure:"-"`
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

	err = viperInstance.Unmarshal(&config)

	if err != nil {
		return
	}

	config.HTTP = HTTP{
		Address: viperInstance.GetString("http.address"),
		Prefork: viperInstance.GetBool("http.prefork"),
		Timeout: viperInstance.GetDuration("http.timeout"),
		BaseURL: viperInstance.GetString("http.baseurl"),
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
	config.FrontendRedirectURL = viperInstance.GetString("frontend_redirect_url")
	return
}
