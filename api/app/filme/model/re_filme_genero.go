package model

import (
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReFilmeGenero struct {
	Id       uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	IdFilme  uuid.UUID `json:"idFilme" gorm:"not null"`
	IdGenero uuid.UUID `json:"idGenero" gorm:"not null"`
}

func (rfg *ReFilmeGenero) BeforeCreate(_ *gorm.DB) (err error) {
	rfg.Id, err = uuid.NewV7()
	return
}

func (rfg *ReFilmeGenero) TableName() string {
	return global.TabelaReFilmeGenero
}
