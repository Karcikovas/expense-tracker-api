package database

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

const ListHardLimit = 100

type Repository[T any] struct {
	DB *gorm.DB
}

func (r Repository[T]) Create(ctx context.Context, tModel *T, transaction *gorm.DB) error {
	result := r.Connection(transaction).WithContext(ctx).Create(tModel)
	if result.Error != nil {
		return fmt.Errorf("%w: %w", ErrUnableToCreate, result.Error)
	}
	return nil
}

func (r Repository[T]) Connection(transaction *gorm.DB) *gorm.DB {
	if transaction == nil {
		return r.DB
	}

	return transaction
}
