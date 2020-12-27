package models

type Subscription struct {
	Model
	Email string `gorm:"not null"`
}
