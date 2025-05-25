package model

import (
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Ingresso struct {
	Id         uuid.UUID `json:"id"`
	IdSessao   uuid.UUID `json:"idSessao"`
	IdUsuario  uuid.UUID `json:"idUsuario"`
	Poltrona   string    `json:"poltrona"`
	CompradoEm time.Time `json:"compradoEm"`
}

func (i *Ingresso) TableName() string {
	return global.TabelaIngresso
}

func (i *Ingresso) BeforeCreate(_ *gorm.DB) (err error) {
	i.Id, err = uuid.NewV7()
	return
}

type CompraIngressoRequest struct {
	IdSessao  uuid.UUID `json:"idSessao" binding:"required"`
	IdUsuario uuid.UUID `json:"idUsuario"`
	Poltronas []string  `json:"poltronas" binding:"required"`
}
