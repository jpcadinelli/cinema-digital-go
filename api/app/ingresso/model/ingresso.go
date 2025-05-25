package model

import (
	"cinema_digital_go/api/app/sessao/model"
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Ingresso struct {
	Id         uuid.UUID    `json:"id"`
	IdSessao   uuid.UUID    `json:"idSessao"`
	Sessao     model.Sessao `json:"sessao" gorm:"foreignKey:id_sessao"`
	IdUsuario  uuid.UUID    `json:"idUsuario"`
	Poltrona   string       `json:"poltrona"`
	CompradoEm time.Time    `json:"compradoEm"`
}

func (i *Ingresso) TableName() string {
	return global.TabelaIngresso
}

func (i *Ingresso) BeforeCreate(_ *gorm.DB) (err error) {
	i.Id, err = uuid.NewV7()
	return
}

func (i Ingresso) ToResponse() IngressoResponse {
	return IngressoResponse{
		Id:         i.Id,
		Filme:      i.Sessao.Filme.Titulo,
		Sala:       i.Sessao.Sala.Nome,
		Poltrona:   i.Poltrona,
		CompradoEm: i.CompradoEm,
	}
}

type CompraIngressoRequest struct {
	IdSessao  uuid.UUID `json:"idSessao" binding:"required"`
	IdUsuario uuid.UUID `json:"idUsuario"`
	Poltronas []string  `json:"poltronas" binding:"required"`
}

type IngressoResponse struct {
	Id         uuid.UUID `json:"id"`
	Filme      string    `json:"filme"`
	Sala       string    `json:"sala"`
	Poltrona   string    `json:"poltrona"`
	CompradoEm time.Time `json:"compradoEm"`
}
