package store

import (
	"context"
	"hw6/internal/models"
)

type Store interface {
	Create(ctx context.Context, laptop *models.Product) error
	All(ctx context.Context) ([]*models.Product, error)
	ByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, laptop *models.Product) error
	Delete(ctx context.Context, id int) error
}
