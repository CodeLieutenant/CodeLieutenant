package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Db *gorm.DB
)

func Connect(dsn string) (err error) {
	config := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent,
				Colorful:      true,
			},
		),
	}

	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), config)

	return err
}
