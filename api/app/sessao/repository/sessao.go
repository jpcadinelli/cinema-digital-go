package repository

import (
	"cinema_digital_go/api/app/sessao/model"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessaoRepository interface {
	Create(sessao *model.Sessao) error
	FindById(id uuid.UUID) (*model.Sessao, error)
	FindAll(preloads ...string) ([]model.Sessao, error)
	Update(sessao *model.Sessao, updateItems map[string]interface{}) (*model.Sessao, error)
	Delete(id uuid.UUID) error
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

func (r *sessaoRepositoryImpl) FindById(id uuid.UUID) (*model.Sessao, error) {
	var sessao model.Sessao

	tx := r.db.First(&sessao, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrSessaoNaoEncontrada
	}

	return &sessao, nil
}

func (r *sessaoRepositoryImpl) FindAll(preloads ...string) ([]model.Sessao, error) {
	var sessoes []model.Sessao

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&sessoes)
	if tx.Error != nil {
		return sessoes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return sessoes, erros.ErrSessaoNaoEncontrada
	}

	return sessoes, nil
}

func (r *sessaoRepositoryImpl) Update(sessao *model.Sessao, updateItems map[string]interface{}) (*model.Sessao, error) {
	tx := r.db.Model(sessao).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrSessaoNaoEncontrada
	}

	return sessao, nil
}

func (r *sessaoRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Sessao{}, "id = ?", id).Error
}
