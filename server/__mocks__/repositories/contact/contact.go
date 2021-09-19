package contact

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Insert(ctx context.Context, contactDto dto.Contact) (models.Contact, error) {
	args := r.Called(ctx, contactDto)

	return args.Get(0).(models.Contact), args.Error(1)
}
