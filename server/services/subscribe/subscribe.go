package subscribe

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leebenson/conform"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/models"
)

type Service interface {
	Subscribe(context.Context, dto.Subscription) (models.Subscription, error)
	Unsubscribe(context.Context, uint64) error
}

type subscriptionService struct {
	db       *pgxpool.Pool
	validate *validator.Validate
}

func (s subscriptionService) Subscribe(ctx context.Context, sub dto.Subscription) (models.Subscription, error) {
	newCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	conn, err := s.db.Acquire(newCtx)

	if err != nil {
		return models.Subscription{}, err
	}
	defer conn.Release()

	if err := conform.Strings(&sub); err != nil {
		return models.Subscription{}, err
	}

	if err := s.validate.Struct(sub); err != nil {
		return models.Subscription{}, err
	}

	tx, err := conn.Begin(newCtx)

	if err != nil {
		return models.Subscription{}, err
	}

	createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()
	sql := `
		INSERT INTO subscriptions(name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	row := tx.QueryRow(
		newCtx,
		sql,
		sub.Name,
		sub.Email,
		createdAt,
		updatedAt,
	)

	var id uint64

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback(newCtx)
		return models.Subscription{}, err
	}

	if err := tx.Commit(newCtx); err != nil {
		_ = tx.Rollback(newCtx)
		return models.Subscription{}, err
	}

	return models.Subscription{
		Model: models.Model{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Name:  sub.Name,
		Email: sub.Email,
	}, nil
}

func (s subscriptionService) Unsubscribe(ctx context.Context, id uint64) error {
	newCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	conn, err := s.db.Acquire(newCtx)

	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.Begin(newCtx)

	if err != nil {
		return err
	}

	sql := "DELETE FROM subscriptions WHERE id = $1;"

	tag, err := tx.Exec(newCtx, sql, id)

	if err != nil {
		_ = tx.Rollback(newCtx)
		return err
	}

	if tag.RowsAffected() != 1 {
		_ = tx.Rollback(newCtx)
		return err
	}

	return tx.Commit(newCtx)
}

func NewSubscriptionService(db *pgxpool.Pool, validate *validator.Validate) Service {
	return subscriptionService{
		db:       db,
		validate: validate,
	}
}
