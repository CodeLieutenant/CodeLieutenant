package config

import (
	"io"

	"github.com/spf13/viper"
)

type (
	DB struct {
		URI      string `json:"uri,omitempty" yaml:"uri,omitempty"`
		Host     string `json:"host,omitempty" yaml:"host,omitempty"`
		User     string `json:"user,omitempty" yaml:"user,omitempty"`
		Password string `json:"password,omitempty" yaml:"password,omitempty"`
		DBName   string `json:"dbname,omitempty" yaml:"dbname,omitempty"`
		TimeZone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
		LogFile  string `json:"logfile,omitempty" yaml:"logfile,omitempty"`
		SSLMode  bool   `json:"sslmode,omitempty" yaml:"sslmode,omitempty"`
		Port     int32  `json:"port,omitempty" yaml:"port,omitempty"`
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

	Config struct {
		Database DB      `json:"postgres" yaml:"postgres"`
		Logging  Logging `json:"logging" yaml:"logging"`
		HTTP     HTTP    `json:"http" yaml:"http"`
		Debug    bool    `json:"debug" yaml:"debug"`
		Locale   string  `json:"locale" yaml:"locale"`
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

	err = viperInstance.UnmarshalKey("postgres", &config.Database)

	return
}
