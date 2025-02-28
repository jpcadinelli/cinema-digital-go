package repository

import (
	"cinema_digital_go/api/app/filme/model"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmeRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*model.Filme, error)
	FindAll(preloads ...string) ([]model.Filme, error)
	Create(filme *model.Filme) error
	Update(filme *model.Filme, updateItems map[string]interface{}) (*model.Filme, error)
	Delete(id uuid.UUID) error
}

type filmeRepositoryImpl struct {
	db *gorm.DB
}

func NewFilmeRepository(db *gorm.DB) FilmeRepository {
	return &filmeRepositoryImpl{db: db}
}

func (r *filmeRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*model.Filme, error) {
	var filme model.Filme

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&filme, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrFilmeNaoEncontrado
	}

	return &filme, nil
}

func (r *filmeRepositoryImpl) FindAll(preloads ...string) ([]model.Filme, error) {
	var filmes []model.Filme

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&filmes)
	if tx.Error != nil {
		return filmes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return filmes, erros.ErrFilmeNaoEncontrado
	}

	return filmes, nil
}

func (r *filmeRepositoryImpl) Create(filme *model.Filme) error {
	return r.db.Create(filme).Error
}

func (r *filmeRepositoryImpl) Update(filme *model.Filme, updateItems map[string]interface{}) (*model.Filme, error) {
	tx := r.db.Model(filme).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrFilmeNaoEncontrado
	}

	return filme, nil
}

func (r *filmeRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Filme{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
