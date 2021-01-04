package models

type Subscription struct {
	Model
	Name  string `gorm:"not null" json:"name,omitempty"`
	Email string `gorm:"not null" json:"email,omitempty"`
}
