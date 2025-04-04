package model

import (
	"cinema_digital_go/api/app/dropdown/model"
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sala struct {
	Id        uuid.UUID `json:"id"`
	Nome      string    `json:"nome" gorm:"unique;not null"`
	Fileiras  string    `json:"fileiras" gorm:"unique;not null"`
	Poltronas int       `json:"poltronas" gorm:"unique;not null"`
}

func (s *Sala) BeforeCreate(_ *gorm.DB) (err error) {
	s.Id, err = uuid.NewV7()
	return
}

func (s *Sala) TableName() string {
	return global.TabelaSala
}

func (s *Sala) SalaToDropdownUUID() *model.DropdownUUID {
	return &model.DropdownUUID{
		Label: s.Nome,
		Value: s.Id,
	}
}
