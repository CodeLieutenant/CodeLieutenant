package services

import "github.com/malusev998/dusanmalusev/models"

type Subscription struct {
	Email string `json:"email" conform:"trim" validate:"required,email,max=150"`
}

type SubscriptionServer interface {
	Subscribe() (models.Subscription, error)
	Unsubscribe() error
}

type subscriptionServer struct{}

func (s subscriptionServer) Subscribe() (models.Subscription, error) {
	panic("implement me")
}

func (s subscriptionServer) Unsubscribe() error {
	panic("implement me")
}
