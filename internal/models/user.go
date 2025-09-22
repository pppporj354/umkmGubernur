package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	FullName     string    `gorm:"size:120;not null"`
	Email        string    `gorm:"size:255;not null;unique"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
