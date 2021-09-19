package contact

import (
	"context"
	"github.com/malusev998/malusev998/server/dto"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/tests"
)

func TestContact(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	pool, drop, err := tests.CreateDatabase()
	assert.NoError(err)
	defer drop()

	repo := &repo{
		db: pool,
	}

	contactDto := dto.Contact{
		Name:    "Test",
		Email:   "test@test.com",
		Subject: "Subject",
		Message: "Test Message",
	}

	model, err := repo.Insert(context.Background(), contactDto)
	if err != nil {
		return
	}

	assert.NoError(err)
	assert.NotEmpty(model)
	assert.GreaterOrEqual(uint64(1), model.ID)

	var count int
	row := pool.QueryRow(context.Background(), "SELECT COUNT(*) as count FROM contacts;")

	assert.NoError(row.Scan(&count))
	assert.Equal(1, count)
}
