package model

import (
	"cinema_digital_go/api/app/dropdown/model"
	"cinema_digital_go/api/pkg/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permissao struct {
	Id        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Descricao string    `json:"descricao"`
}

func (p *Permissao) BeforeCreate(_ *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	p.Id = id
	return nil
}

func (p *Permissao) TableName() string {
	return global.TabelaPermissao
}

func (p *Permissao) PermissaoToDropdownUUID() *model.DropdownUUID {
	return &model.DropdownUUID{
		Label: p.Nome,
		Value: p.Id,
	}
}

type PermissaoUsuario struct {
	Id          uuid.UUID `json:"id"`
	IdPermissao uuid.UUID `json:"idPermissao" gorm:"column:id_permissao"`
	IdUsuario   uuid.UUID `json:"idUsuario" gorm:"column:id_usuario"`
}

func (pu *PermissaoUsuario) BeforeCreate(_ *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	pu.Id = id
	return nil
}

func (pu *PermissaoUsuario) TableName() string {
	return global.TabelaPermissaoUsuario
}
