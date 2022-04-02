package repository

import (
	"context"

	"github.com/yeldisbayev/thrid-group/library/internal/domain"
)

type BookRepo interface {
	Store(ctx context.Context, book domain.Book) (*domain.Book, error)
	Get(ctx context.Context, id string) (*domain.Book, error)
	Remove(ctx context.Context, id string) error
}
