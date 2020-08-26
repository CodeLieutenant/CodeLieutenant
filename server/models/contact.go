package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null"`
	Subject string `gorm:"not null"`
	Message string `gorm:"not null"`
}
