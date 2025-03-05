package model

import (
	"cinema_digital_go/api/app/dropdown/model"
	"cinema_digital_go/api/pkg/global"
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Genero struct {
	Id        uuid.UUID `json:"id"`
	Nome      string    `json:"nome" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
}

func (g *Genero) BeforeCreate(_ *gorm.DB) (err error) {
	g.Id, err = uuid.NewV7()
	g.CreatedAt = time.Now()
	return
}

func (g *Genero) TableName() string {
	return global.TabelaGenero
}

func (g *Genero) GeneroToDropdownUUID() *model.DropdownUUID {
	return &model.DropdownUUID{
		Label: g.Nome,
		Value: g.Id,
	}
}
