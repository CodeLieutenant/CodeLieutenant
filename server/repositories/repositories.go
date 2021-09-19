package repositories

import (
	"context"
	"github.com/malusev998/malusev998/server/models"

	"github.com/malusev998/malusev998/server/dto"
)

type (
	Contact interface {
		Insert(ctx context.Context, contactDto dto.Contact) (models.Contact, error)
	}
	Post interface {
		Paginate(ctx context.Context, lastId uint64, perPage uint64) ([]models.Post, error)
		GetOne(ctx context.Context, id uint64) (models.Post, error)
		Delete(ctx context.Context, id uint64) error
	}
	Subscribe interface {
		Insert(context.Context, dto.Subscription) (models.Subscription, error)
		Remove(ctx context.Context, id uint64) error
	}
)
