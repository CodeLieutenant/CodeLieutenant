package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Email string `gorm:"not null"`
}
