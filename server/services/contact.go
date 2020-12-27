package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
	"gorm.io/gorm"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/models"
)

type ContactService interface {
	AddMessage(dto.Contact) (models.Contact, error)
}

type contactService struct {
	db        *gorm.DB
	validator *validator.Validate
}

func (c contactService) AddMessage(contactDto dto.Contact) (models.Contact, error) {
	if err := conform.Strings(&contactDto); err != nil {
		return models.Contact{}, err
	}

	if err := c.validator.Struct(c); err != nil {
		return models.Contact{}, err
	}

	contact := models.Contact{
		Name:    contactDto.Name,
		Email:   contactDto.Email,
		Subject: contactDto.Subject,
		Message: contactDto.Message,
	}

	result := c.db.Create(&contact)

	if result.Error != nil {
		return models.Contact{}, result.Error
	}

	return contact, nil
}

func NewContactService(db *gorm.DB, validate *validator.Validate) ContactService {
	return contactService{
		db:        db,
		validator: validate,
	}
}
