package repository

import (
	"cinema_digital_go/api/app/genero/model"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneroRepository interface {
	Create(genero *model.Genero) error
	FindById(id uuid.UUID) (*model.Genero, error)
	FindAll(preloads ...string) ([]model.Genero, error)
}

type generoRepositoryImpl struct {
	db *gorm.DB
}

func NewGeneroRepository(db *gorm.DB) GeneroRepository {
	return &generoRepositoryImpl{db: db}
}

func (r *generoRepositoryImpl) Create(genero *model.Genero) error {
	return r.db.Create(genero).Error
}

func (r *generoRepositoryImpl) FindById(id uuid.UUID) (*model.Genero, error) {
	var genero model.Genero

	tx := r.db.First(&genero, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrGeneroNaoEncontrado
	}

	return &genero, nil
}

func (r *generoRepositoryImpl) FindAll(preloads ...string) ([]model.Genero, error) {
	var generos []model.Genero

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&generos)
	if tx.Error != nil {
		return generos, tx.Error
	}
	if tx.RowsAffected == 0 {
		return generos, erros.ErrGeneroNaoEncontrado
	}

	return generos, nil
}
