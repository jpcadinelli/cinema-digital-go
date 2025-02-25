package repository

import (
	"cinema_digital_go/api/app/genero/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneroRepository interface {
	Create(genero *model.Genero) error
	GetAll() ([]model.Genero, error)
	Update(genero *model.Genero) error
	Delete(id uuid.UUID) error
	GetByName(nome string) (*model.Genero, error)
}

type generoRepository struct {
	db *gorm.DB
}

func NewGeneroRepository(db *gorm.DB) GeneroRepository {
	return &generoRepository{db}
}

func (r *generoRepository) Create(genero *model.Genero) error {
	if genero.ID == uuid.Nil {
		genero.ID = uuid.New()
	}

	return r.db.Table("genero").Create(genero).Error
}

func (r *generoRepository) GetAll() ([]model.Genero, error) {
	var generos []model.Genero
	err := r.db.Table("genero").Find(&generos).Error
	return generos, err
}

func (r *generoRepository) Update(genero *model.Genero) error {
	return r.db.Table("genero").Save(genero).Error
}

func (r *generoRepository) Delete(id uuid.UUID) error {
	return r.db.Table("genero").Where("id = ?", id).Delete(&model.Genero{}).Error
}

func (r *generoRepository) GetByName(nome string) (*model.Genero, error) {
	var genero model.Genero
	err := r.db.Table("genero").Where("nome = ?", nome).First(&genero).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &genero, nil
}
