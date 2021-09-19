package services

import (
	"context"
	"io"

	"github.com/jordan-wright/email"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

type (
	PostService interface {
		Get(context.Context, uint64, uint64) ([]models.Post, error)
		GetOne(context.Context, uint64) (models.Post, error)
		Delete(context.Context, uint64) error
	}

	ContactService interface {
		AddMessage(ctx context.Context, contactDto dto.Contact) (models.Contact, error)
	}

	SubscribeService interface {
		Subscribe(context.Context, dto.Subscription) (models.Subscription, error)
		Unsubscribe(context.Context, uint64) error
	}

	EmailService interface {
		io.Closer
		NewEmail() *email.Email
		Send(to []string, subject string, body []byte) error
		SendEmail(*email.Email) error
	}
)
