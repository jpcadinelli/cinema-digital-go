package model

import (
	"cinema_digital_go/api/pkg/global"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Sessao struct {
	Id              uuid.UUID `json:"id"`
	IdFilme         uuid.UUID `json:"idFilme"`
	IdSala          uuid.UUID `json:"idSala"`
	DataInicio      time.Time `json:"dataInicio"`
	DataFim         time.Time `json:"dataFim"`
	PrecoIngresso   float64   `json:"precoIngresso"`
	Disponibilidade int       `json:"disponibilidade"`
}

func (s *Sessao) TableName() string {
	return global.TabelaSessao
}

func (s *Sessao) BeforeCreate(_ *gorm.DB) (err error) {
	s.Id, err = uuid.NewV7()
	return
}

func (s *Sessao) Validar() error {
	if s.IdFilme == uuid.Nil {
		return errors.New("idFilme is required")
	}
	if s.IdSala == uuid.Nil {
		return errors.New("idSala is required")
	}
	if s.DataInicio == (time.Time{}) {
		return errors.New("dataInicio is required")
	}
	if s.DataFim == (time.Time{}) {
		return errors.New("dataFim is required")
	}
	if s.PrecoIngresso == 0 {
		return errors.New("precoIngresso is required")
	}
	if s.Disponibilidade == 0 {
		return errors.New("disponibilidade is required")
	}
	return nil
}
