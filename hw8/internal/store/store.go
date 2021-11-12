package store

import (
	"context"
	"hw8/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Anime() TitleRepository
	Manga() TitleRepository
	Ranobe() TitleRepository
}

type TitleRepository interface {
	Create(ctx context.Context, category *models.Title) error
	All(ctx context.Context) ([]*models.Title, error)
	ByID(ctx context.Context, id int) (*models.Title, error)
	Update(ctx context.Context, category *models.Title) error
	Delete(ctx context.Context, id int) error
}