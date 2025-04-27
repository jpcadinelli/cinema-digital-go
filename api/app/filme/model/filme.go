package model

import (
	modelGen "cinema_digital_go/api/app/genero/model"
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Filme struct {
	Id                uuid.UUID         `json:"id"`
	Titulo            string            `json:"titulo"`
	Sinopse           string            `json:"sinopse"`
	Diretor           string            `json:"diretor"`
	Duracao           uint              `json:"duracao"`
	AnoLancamento     time.Time         `json:"anoLancamento"`
	Classificacao     uint              `json:"classificacao"`
	Nota              float64           `json:"nota"`
	Criado            time.Time         `json:"criado"`
	Atualizado        time.Time         `json:"atualizado"`
	Excluido          *gorm.DeletedAt   `json:"excluido"`
	IdUsuarioRegistro uuid.UUID         `json:"idUsuarioRegistro"`
	Generos           []modelGen.Genero `json:"generos" gorm:"many2many:re_filme_genero;foreignKey:id;References:id;re_filme_genero;joinForeignKey:id_filme;joinReferences:id_genero"`
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
	return global.TabelaFilme
}

func (f *Filme) GetOnlyFilme() *Filme {
	return &Filme{
		Id:                f.Id,
		Titulo:            f.Titulo,
		Sinopse:           f.Sinopse,
		Diretor:           f.Diretor,
		Duracao:           f.Duracao,
		AnoLancamento:     f.AnoLancamento,
		Classificacao:     f.Classificacao,
		Nota:              f.Nota,
		Criado:            f.Criado,
		Atualizado:        f.Atualizado,
		Excluido:          f.Excluido,
		IdUsuarioRegistro: f.IdUsuarioRegistro,
	}
}
