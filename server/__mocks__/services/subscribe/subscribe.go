package subscribe

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

type (
	ServiceMock struct {
		mock.Mock
	}
)

func (c *ServiceMock) Subscribe(ctx context.Context, subscribeDto dto.Subscription) (models.Subscription, error) {
	args := c.Called(ctx, subscribeDto)

	return args.Get(0).(models.Subscription), args.Error(1)
}

func (c *ServiceMock) Unsubscribe(ctx context.Context, id uint64) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}
