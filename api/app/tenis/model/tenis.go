package model

import (
	"github.com/google/uuid"
	"time"
)

type Tenis struct {
	Id        uuid.UUID `json:"id"`
	Marca     string    `json:"marca" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
}
