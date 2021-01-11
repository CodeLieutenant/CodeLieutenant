package subscribe

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/models"
	"github.com/malusev998/dusanmalusev/services/email"
)

type subscriptionWithEmail struct {
	service      Service
	emailService email.Interface
}

func (s subscriptionWithEmail) Subscribe(ctx context.Context, sub dto.Subscription) (models.Subscription, error) {
	// TODO Send E-MAIL

	return s.service.Subscribe(ctx, sub)
}

func (s subscriptionWithEmail) Unsubscribe(ctx context.Context, id uint64) error {
	// TODO: Send EMAIL
	return s.service.Unsubscribe(ctx, id)
}

func NewSubscriptionWithEmail(email email.Interface, db *pgxpool.Pool, validate *validator.Validate) Service {
	return subscriptionWithEmail{
		service:      NewSubscriptionService(db, validate),
		emailService: email,
	}
}
