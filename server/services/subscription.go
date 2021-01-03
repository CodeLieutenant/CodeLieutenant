package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/models"
)

type SubscriptionService interface {
	Subscribe(dto.Subscription) (models.Subscription, error)
	Unsubscribe(uint) error
}

type subscriptionService struct {
	db       *pgxpool.Pool
	validate *validator.Validate
}

func (s subscriptionService) Subscribe(sub dto.Subscription) (models.Subscription, error) {
	panic("implement me")
}

func (s subscriptionService) Unsubscribe(id uint) error {
	panic("implement me")
}

func NewSubscriptionService(db *pgxpool.Pool, validate *validator.Validate) SubscriptionService {
	return subscriptionService{
		db:       db,
		validate: validate,
	}
}
