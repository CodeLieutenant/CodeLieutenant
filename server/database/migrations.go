package database

import (
	"github.com/malusev998/dusanmalusev/models"
)

func RunMigrations() error {
	m := []interface{}{
		models.Contact{},
		models.Subscription{},
	}


	return Db.AutoMigrate(m...)
}
