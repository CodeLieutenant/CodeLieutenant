package services

import (
	"github.com/malusev998/dusanmalusev/database"
	"github.com/malusev998/dusanmalusev/models"
)

type Contact struct {
	Name    string `json:"name" conform:"trim" validate:"required,alphanumericunicodespace,max=50"`
	Email   string `json:"email" conform:"trim" validate:"required,email,max=150"`
	Subject string `json:"subject" conform:"trim" validate:"required,alphanumericunicodespace,min=3,max=150"`
	Message string `json:"message" conform:"trim" validate:"required,min=3,max=1000"`
}

type ContactService interface {
	AddMessage(Contact) (models.Contact, error)
}

type contactService struct{}

func (c contactService) AddMessage(contactDto Contact) (models.Contact, error) {
	contact := models.Contact{
		Name:    contactDto.Name,
		Email:   contactDto.Email,
		Subject: contactDto.Subject,
		Message: contactDto.Message,
	}

	result := database.Db.Create(&contact)

	if result.Error != nil {
		return models.Contact{}, result.Error
	}

	return contact, nil
}

func NewContactService() ContactService {
	return contactService{}
}
