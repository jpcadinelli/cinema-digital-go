package repository

import (
	"cinema_digital_go/api/app/sessao/model"
	"gorm.io/gorm"
)

type SessaoRepository interface {
	Create(sessao *model.Sessao) error
}

type sessaoRepositoryImpl struct {
	db *gorm.DB
}

func NewSessaoRepository(db *gorm.DB) SessaoRepository {
	return &sessaoRepositoryImpl{db: db}
}

func (r *sessaoRepositoryImpl) Create(sessao *model.Sessao) error {
	return r.db.Model(&model.Sessao{}).Create(sessao).Error
}
