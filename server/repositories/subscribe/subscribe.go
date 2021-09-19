package subscribe

import (
	"context"
	"github.com/malusev998/malusev998/server/database"
	"github.com/malusev998/malusev998/server/repositories"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

const (
	insertStatement = `
		INSERT INTO subscriptions(name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	deleteStatement = "DELETE FROM subscriptions WHERE id = $1;"
)

type repo struct {
	db       *pgxpool.Pool
}

func (s repo) Insert(ctx context.Context, sub dto.Subscription) (models.Subscription, error) {
	conn, err := s.db.Acquire(ctx)

	if err != nil {
		return models.Subscription{}, err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	if err != nil {
		return models.Subscription{}, err
	}

	createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()

	row := tx.QueryRow(
		ctx,
		insertStatement,
		sub.Name,
		sub.Email,
		createdAt,
		updatedAt,
	)

	var id uint64

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback(ctx)
		return models.Subscription{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
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

func (s repo) Remove(ctx context.Context, id uint64) error {
	conn, err := s.db.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.Begin(ctx)

	if err != nil {
		return err
	}

	tag, err := tx.Exec(ctx, deleteStatement, id)

	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if tag.RowsAffected() != 1 {
		_ = tx.Rollback(ctx)
		return database.ErrNotFound
	}

	return tx.Commit(ctx)
}

func New(db *pgxpool.Pool) repositories.Subscribe {
	return repo{
		db:       db,
	}
}
