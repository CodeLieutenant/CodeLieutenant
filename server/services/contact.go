package services

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leebenson/conform"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/models"
)

type ContactService interface {
	AddMessage(ctx context.Context, contactDto dto.Contact) (models.Contact, error)
}

type contactService struct {
	db        *pgxpool.Pool
	validate *validator.Validate
}

func (c contactService) AddMessage(ctx context.Context, contactDto dto.Contact) (models.Contact, error) {
	newCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	conn, err := c.db.Acquire(newCtx)

	if err != nil {
		return models.Contact{}, err
	}

	defer conn.Release()

	if err := conform.Strings(&contactDto); err != nil {
		return models.Contact{}, err
	}

	if err := c.validate.Struct(contactDto); err != nil {
		return models.Contact{}, err
	}

	tx, err := conn.Begin(newCtx)

	if err != nil {
		return models.Contact{}, err
	}

	createdAt, updatedAt := time.Now().UTC(), time.Now().UTC()
	sql := `
		INSERT INTO contacts(name, email, subject, message, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	row := tx.QueryRow(
		newCtx,
		sql,
		contactDto.Name,
		contactDto.Email,
		contactDto.Subject,
		contactDto.Message,
		createdAt,
		updatedAt,
	)

	var id uint64

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback(newCtx)
		return models.Contact{}, err
	}

	if err := tx.Commit(newCtx); err != nil {
		_ = tx.Rollback(newCtx)
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

func NewContactService(db *pgxpool.Pool, validate *validator.Validate) ContactService {
	return contactService{
		db:        db,
		validate: validate,
	}
}
