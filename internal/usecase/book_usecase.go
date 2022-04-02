package usecase

import (
	"context"

	"github.com/yeldisbayev/thrid-group/library/internal/domain"
)

type BookUsecase interface {
	Create(ctx context.Context, book domain.Book) (*domain.Book, error)
	Get(ctx context.Context, id string) (*domain.Book, error)
	Delete(ctx context.Context, id string) error
}
