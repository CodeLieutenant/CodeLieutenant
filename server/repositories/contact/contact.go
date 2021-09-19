package contact

import (
	"context"
	"github.com/malusev998/malusev998/server/repositories"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/models"
)

const (
	insertStatement = `
		INSERT INTO contacts(name, email, subject, message, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
)

type repo struct {
	db *pgxpool.Pool
}

func (c repo) Insert(ctx context.Context, contactDto dto.Contact) (models.Contact, error) {
	conn, err := c.db.Acquire(ctx)

	if err != nil {
		return models.Contact{}, err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	if err != nil {
		return models.Contact{}, err
	}

	createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()

	row := tx.QueryRow(
		ctx,
		insertStatement,
		contactDto.Name,
		contactDto.Email,
		contactDto.Subject,
		contactDto.Message,
		createdAt,
		updatedAt,
	)

	var id uint64

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback(ctx)
		return models.Contact{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		return models.Contact{}, err
	}

	return models.Contact{
		Model: models.Model{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Name:    contactDto.Name,
		Email:   contactDto.Email,
		Subject: contactDto.Subject,
		Message: contactDto.Message,
	}, nil
}

func New(db *pgxpool.Pool) repositories.Contact {
	return repo{
		db: db,
	}
}
