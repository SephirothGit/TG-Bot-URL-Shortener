package repository

import (
	"API/internal/models"
	"context"
	"database/sql"
)

type LinkRepository interface {
	Create(ctx context.Context, link *models.Link) error
	GetAll(ctx context.Context) ([]*models.Link, error)
	GetByCode(ctx context.Context, code string) (*models.Link, error)
	Delete(ctx context.Context, code string) error
	IncrementClicks(ctx context.Context, code string) error
}

type PostgresLinkRepo struct {
	db *sql.DB
}

func NewPostgresLinkRepo(db *sql.DB) LinkRepository {
	return &PostgresLinkRepo{db: db}
}

func (r *PostgresLinkRepo) Create(ctx context.Context, link *models.Link) error {
	err := r.db.QueryRowContext(ctx, "INSERT INTO links (original_url, short_code, clicks, created_at) VALUES ($1, $2, 0, NOW()) RETURNING id, created_at",
		link.OriginalURL, link.ShortCode).Scan(&link.ID, &link.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresLinkRepo) GetAll(ctx context.Context) ([]*models.Link, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, original_url, short_code, clicks, created_at FROM links")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*models.Link

	for rows.Next() {
		link := &models.Link{}
		if err := rows.Scan(&link.ID, &link.OriginalURL, &link.ShortCode, &link.Clicks, &link.CreatedAt); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (r *PostgresLinkRepo) GetByCode(ctx context.Context, code string) (*models.Link, error) {

	link := &models.Link{}

	err := r.db.QueryRowContext(ctx, "SELECT id, original_url, short_code, clicks, created_at FROM links WHERE short_code = $1", code).Scan(&link.ID, &link.OriginalURL, &link.ShortCode, &link.Clicks, &link.CreatedAt)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (r *PostgresLinkRepo) Delete(ctx context.Context, code string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM links WHERE short_code = $1", code)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresLinkRepo) IncrementClicks(ctx context.Context, code string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE links SET clicks = clicks + 1 WHERE short_code = $1", code)
	if err != nil {
		return err
	}
	return nil
}
