package contact

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
	"github.com/malusev998/malusev998/server/repositories"
	"github.com/malusev998/malusev998/server/services"
)

type service struct {
	repo     repositories.Contact
	validate *validator.Validate
}

func New(repo repositories.Contact, validate *validator.Validate) services.ContactService {
	return service{
		repo:     repo,
		validate: validate,
	}
}

func (s service) AddMessage(ctx context.Context, contactDto dto.Contact) (models.Contact, error) {
	newCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	if err := conform.Strings(&contactDto); err != nil {
		return models.Contact{}, err
	}

	if err := s.validate.Struct(contactDto); err != nil {
		return models.Contact{}, err
	}

	return s.repo.Insert(newCtx, contactDto)
}
