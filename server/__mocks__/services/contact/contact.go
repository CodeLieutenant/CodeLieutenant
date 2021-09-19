package contact

import (
	"context"
	"github.com/stretchr/testify/mock"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

type ServiceMock struct {
	mock.Mock
}

func (c *ServiceMock) AddMessage(ctx context.Context, contactDto dto.Contact) (models.Contact, error) {
	args := c.Called(ctx, contactDto)

	return args.Get(0).(models.Contact), args.Error(1)
}
