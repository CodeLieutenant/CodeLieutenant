package subscribe

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/tests"
)

type SubscribeServiceTest struct {
	tests.DBTestCase
}

func (s *SubscribeServiceTest) TestSubscribe_ValidationError() {
	subDto := dto.Subscription{
		Name:  "test",
		Email: "test@.com",
	}

	service := subscriptionService{
		db:       s.DB,
		validate: tests.GetValidator(),
	}

	sub, err := service.Subscribe(context.Background(), subDto)

	s.Error(err)
	s.Empty(sub)

	row := s.DB.QueryRow(s.Ctx, "SELECT COUNT(id) FROM subscriptions;")
	var count uint64
	s.NoError(row.Scan(&count))
	s.Equal(uint64(0), count)
}

func (s *SubscribeServiceTest) TestSubscribe_Success() {
	subDto := dto.Subscription{
		Name:  "test",
		Email: "test@test.com",
	}

	service := subscriptionService{
		db:       s.DB,
		validate: tests.GetValidator(),
	}

	sub, err := service.Subscribe(context.Background(), subDto)

	s.NoError(err)
	s.NotEmpty(sub)

	s.NotEqual(0, sub.ID)
	s.Equal("test", sub.Name)
	s.Equal("test@test.com", sub.Email)
	row := s.DB.QueryRow(s.Ctx, "SELECT COUNT(id) FROM subscriptions;")
	var count uint64
	s.NoError(row.Scan(&count))
	s.Equal(uint64(1), count)
}

func (s *SubscribeServiceTest) TestUnsubscribe_Success() {
	createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()
	sql := `
		INSERT INTO subscriptions(name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	db, _ := s.DB.Acquire(s.Ctx)
	row := s.DB.QueryRow(
		s.Ctx,
		sql,
		"test",
		"test@test.com",
		createdAt,
		updatedAt,
	)

	db.Release()

	var count uint64
	var id uint64

	s.NoError(row.Scan(&id))

	service := subscriptionService{
		db:       s.DB,
		validate: tests.GetValidator(),
	}

	err := service.Unsubscribe(s.Ctx, id)
	s.NoError(err)

	row = s.DB.QueryRow(s.Ctx, "SELECT COUNT(id) FROM subscriptions WHERE id = $1;", id)
	s.NoError(row.Scan(&count))
	s.Equal(uint64(0), count)
}

func (s *SubscribeServiceTest) TearDownTest() {
	_, err := s.DB.Exec(s.Ctx, "DELETE FROM subscriptions;")
	if err != nil {
		s.FailNow("Failed cleaning up the table", err.Error())
	}
}

func TestSubscribeService(t *testing.T) {
	t.Parallel()
	suite.Run(t, &SubscribeServiceTest{
		DBTestCase: tests.DBTestCase{
			Ctx:              context.Background(),
			ConnectionString: tests.GetConnectionString(),
		},
	})
}
