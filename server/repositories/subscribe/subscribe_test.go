package subscribe

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/tests"
)

func TestSubscribe_AlreadyExists(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	name := "Test User"
	email := "test@test.com"

	pool, drop, err := tests.CreateDatabase()
	assert.NoError(err)
	defer drop()

	repo := &repo{
		db: pool,
	}

	createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()

	_, err = pool.Exec(
		context.Background(),
		insertStatement,
		name,
		email,
		createdAt,
		updatedAt,
	)
	assert.NoError(err)

	contactDto := dto.Subscription{
		Name:  name,
		Email: email,
	}

	subscription, err := repo.Insert(context.Background(), contactDto)

	assert.Error(err)
	assert.Contains(err.Error(), "duplicate key value violates unique constraint")
	assert.Empty(subscription)
}

func TestSubscribeSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	pool, drop, err := tests.CreateDatabase()
	assert.NoError(err)
	defer drop()

	repo := &repo{
		db: pool,
	}

	contactDto := dto.Subscription{
		Name:  "Test User",
		Email: "test@test.com",
	}

	subscription, err := repo.Insert(context.Background(), contactDto)

	assert.NoError(err)
	assert.NotEmpty(subscription)
	assert.GreaterOrEqual(uint64(1), subscription.ID)

	row := pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM subscriptions")
	var count uint64
	assert.NoError(row.Scan(&count))

	assert.Equal(uint64(1), count)
}
