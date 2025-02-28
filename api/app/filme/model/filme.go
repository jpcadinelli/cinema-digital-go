package model

import (
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Filme struct {
	Id                uuid.UUID       `json:"id"`
	Titulo            string          `json:"titulo"`
	Sinopse           string          `json:"sinopse"`
	Diretor           string          `json:"diretor"`
	Duracao           uint            `json:"duracao"`
	AnoLancamento     time.Time       `json:"ano_lancamento"`
	Classificacao     uint            `json:"classificacao"`
	Nota              float64         `json:"nota"`
	Criado            time.Time       `json:"criado"`
	Atualizado        time.Time       `json:"atualizado"`
	Excluido          *gorm.DeletedAt `json:"excluido"`
	IdUsuarioRegistro uuid.UUID       `json:"id_usuario_registro"`
}

func (f *Filme) BeforeCreate(_ *gorm.DB) (err error) {
	f.Id, err = uuid.NewV7()
	f.Criado = time.Now()
	return
}

func (f *Filme) BeforeUpdate(tx *gorm.DB) (err error) {
	f.Atualizado = time.Now()
	tx.Statement.SetColumn("atualizado", f.Atualizado)
	return
}

func (f *Filme) TableName() string {
	return global.TableFilme
}
