package database

import (
	"github.com/malusev998/dusanmalusev/models"
)

func RunMigrations() error {
	return Db.AutoMigrate(&models.Contact{}, &models.Subscription{})
}
