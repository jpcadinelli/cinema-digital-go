package model

import (
	"time"

	"github.com/google/uuid"
)

type Genero struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
}
