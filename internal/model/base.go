package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        []byte `gorm:"type:binary(16);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseModel) BeforeCrete(_ *gorm.DB) error {
	u, err := uuid.NewV7()

	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnableToGenerateUUID, err)
	}

	b.ID = u[:]

	return nil
}
