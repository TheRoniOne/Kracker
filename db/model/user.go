package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:NewV7();primaryKey;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
