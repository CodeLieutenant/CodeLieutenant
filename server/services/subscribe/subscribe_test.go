package subscribe

import (
	"context"
	"errors"
	"testing"

	"github.com/malusev998/malusev998/server/__mocks__/repositories/subscribe"
	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
	"github.com/malusev998/malusev998/server/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSubscribe_ValidationError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	subDto := dto.Subscription{
		Name:  "test",
		Email: "test@.com",
	}

	repo := new(subscribe.RepositoryMock)

	s := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	sub, err := s.Subscribe(context.Background(), subDto)

	assert.Error(err)
	assert.Empty(sub)

	repo.AssertNotCalled(t, "Insert")

	//row := s.DB.QueryRow(s.Ctx, "SELECT COUNT(id) FROM subscriptions;")
	//var count uint64
	//s.NoError(row.Scan(&count))
	//s.Equal(uint64(0), count)
}

func TestSubscribe_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	subDto := dto.Subscription{
		Name:  "test",
		Email: "test@test.com",
	}

	subscription := models.Subscription{
		Model: models.Model{
			ID: 1,
		},
		Name:  subDto.Name,
		Email: subDto.Email,
	}

	repo := new(subscribe.RepositoryMock)

	s := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	repo.On("Insert", mock.Anything, subDto).
		Once().
		Return(subscription, nil)

	sub, err := s.Subscribe(context.Background(), subDto)

	assert.NoError(err)
	assert.NotEmpty(sub)
	assert.Equal(uint64(1), subscription.ID)
	assert.Equal(subDto.Name, subscription.Name)
	assert.Equal(subDto.Email, subscription.Email)

	repo.AssertExpectations(t)
	//s.NotEqual(0, sub.ID)
	//s.Equal("test", sub.Name)
	//s.Equal("test@test.com", sub.Email)
	//row := s.DB.QueryRow(s.Ctx, "SELECT COUNT(id) FROM subscriptions;")
	//var count uint64
	//s.NoError(row.Scan(&count))
	//s.Equal(uint64(1), count)
}

func TestUnsubscribe_DeleteFailed(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	id := uint64(1)

	repo := new(subscribe.RepositoryMock)

	service := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	repo.On("Remove", mock.Anything, id).
		Once().
		Return(errors.New("remove failed"))

	err := service.Unsubscribe(context.Background(), id)
	assert.Error(err)
	assert.EqualError(err, "remove failed")

	repo.AssertExpectations(t)
}

func TestUnsubscribe_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	//createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()
	//sql := `
	//	INSERT INTO subscriptions(name, email, created_at, updated_at)
	//	VALUES ($1, $2, $3, $4)
	//	RETURNING id;
	//`
	//db, _ := s.DB.Acquire(s.Ctx)
	//row := s.DB.QueryRow(
	//	s.Ctx,
	//	sql,
	//	"test",
	//	"test@test.com",
	//	createdAt,
	//	updatedAt,
	//)
	//
	//db.Release()

	//var count uint64
	//var id uint64

	//s.NoError(row.Scan(&id))

	id := uint64(1)

	repo := new(subscribe.RepositoryMock)

	service := service{
		repo:     repo,
		validate: tests.GetValidator(),
	}

	repo.On("Remove", mock.Anything, id).
		Once().
		Return(nil)

	err := service.Unsubscribe(context.Background(), id)
	assert.NoError(err)

	repo.AssertExpectations(t)

	//row = s.DB.QueryRow(s.Ctx, "SELECT COUNT(id) FROM subscriptions WHERE id = $1;", id)
	//s.NoError(row.Scan(&count))
	//s.Equal(uint64(0), count)
}
