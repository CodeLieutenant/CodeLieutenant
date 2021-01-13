package models

type Contact struct {
	Model
	Name    string `gorm:"not null" json:"name,omitempty" bson:"name"`
	Email   string `gorm:"not null" json:"email,omitempty" bson:"email"`
	Subject string `gorm:"not null" json:"subject,omitempty" bson:"subject"`
	Message string `gorm:"not null" json:"message,omitempty" bson:"message"`
}
