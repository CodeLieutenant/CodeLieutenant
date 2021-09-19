package subscribe

import (
	"context"
	"github.com/malusev998/malusev998/server/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
	"github.com/malusev998/malusev998/server/repositories"
)

type service struct {
	repo     repositories.Subscribe
	validate *validator.Validate
}

func (s service) Subscribe(ctx context.Context, subscription dto.Subscription) (models.Subscription, error) {
	newCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	if err := conform.Strings(&subscription); err != nil {
		return models.Subscription{}, err
	}

	if err := s.validate.Struct(subscription); err != nil {
		return models.Subscription{}, err
	}

	return s.repo.Insert(newCtx, subscription)
}

func (s service) Unsubscribe(ctx context.Context, id uint64) error {
	newCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	// TODO: send email

	return s.repo.Remove(newCtx, id)
}

func New(repo repositories.Subscribe, validate *validator.Validate) services.SubscribeService {
	return service{
		repo:     repo,
		validate: validate,
	}
}
