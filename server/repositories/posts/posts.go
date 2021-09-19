package posts

import (
	"context"
	"errors"
	"github.com/malusev998/malusev998/server/repositories"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/malusev998/malusev998/server/models"
)

const (
	paginateStatement = `
		SELECT 
			p.id, 
			p.title, 
			p.slug, 
			p.content, 
			p.created_at, 
			p.updated_at 
		FROM posts p 
		WHERE p.id > $1 
		ORDER BY p.created_at DESC 
		LIMIT $2;
	`

	getOneStatement = `
		SELECT 
			p.id, 
			p.title, 
			p.slug, 
			p.content, 
			p.created_at, 
			p.updated_at
		FROM posts p
		WHERE p.id = $1;
	`

	deleteStatement = `DELETE FROM posts p WHERE p.id = $1;`
)

var (
	ErrDeleteFailed = errors.New("delete failed")
)

type (
	repo struct {
		db *pgxpool.Pool
	}
)

func New(db *pgxpool.Pool) repositories.Post {
	return repo{
		db: db,
	}
}

func (s repo) Paginate(ctx context.Context, lastId uint64, perPage uint64) ([]models.Post, error) {
	conn, err := s.db.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, paginateStatement, lastId, perPage)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]models.Post, 0, perPage)

	for rows.Next() {
		post := models.Post{}

		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (s repo) GetOne(ctx context.Context, id uint64) (models.Post, error) {
	conn, err := s.db.Acquire(ctx)

	if err != nil {
		return models.Post{}, err
	}

	defer conn.Release()

	rows, err := conn.Query(ctx, getOneStatement, id)

	if err != nil {
		return models.Post{}, err
	}

	defer rows.Close()

	post := models.Post{}

	err = rows.Scan(
		&post.ID,
		&post.Title,
		&post.Slug,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (s repo) Delete(ctx context.Context, id uint64) error {
	conn, err := s.db.Acquire(ctx)

	if err != nil {
		return err
	}

	defer conn.Release()

	tag, err := conn.Exec(ctx, deleteStatement, id)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return ErrDeleteFailed
	}

	return nil
}
