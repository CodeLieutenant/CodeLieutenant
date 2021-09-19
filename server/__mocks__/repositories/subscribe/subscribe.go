package subscribe

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Insert(ctx context.Context, subscription dto.Subscription) (models.Subscription, error) {
	args := r.Called(ctx, subscription)

	return args.Get(0).(models.Subscription), args.Error(1)
}

func (r *RepositoryMock) Remove(ctx context.Context, id uint64) error {
	args := r.Called(ctx, id)

	return args.Error(0)
}
