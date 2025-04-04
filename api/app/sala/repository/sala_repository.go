package repository

import (
	"cinema_digital_go/api/app/sala/model"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalaRepository interface {
	Create(sala *model.Sala) error
	FindById(id uuid.UUID) (*model.Sala, error)
	FindAll(preloads ...string) ([]model.Sala, error)
	Update(sala *model.Sala, updateItems map[string]interface{}) (*model.Sala, error)
	Delete(id uuid.UUID) error
}

type salaRepositoryImpl struct {
	db *gorm.DB
}

func NewSalaRepository(db *gorm.DB) SalaRepository {
	return &salaRepositoryImpl{db: db}
}

func (r *salaRepositoryImpl) Create(sala *model.Sala) error {
	return r.db.Create(sala).Error
}

func (r *salaRepositoryImpl) FindById(id uuid.UUID) (*model.Sala, error) {
	var sala model.Sala

	tx := r.db.First(&sala, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrSalaNaoEncontrada
	}

	return &sala, nil
}

func (r *salaRepositoryImpl) FindAll(preloads ...string) ([]model.Sala, error) {
	var salas []model.Sala

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&salas)
	if tx.Error != nil {
		return salas, tx.Error
	}
	if tx.RowsAffected == 0 {
		return salas, erros.ErrSalaNaoEncontrada
	}

	return salas, nil
}

func (r *salaRepositoryImpl) Update(sala *model.Sala, updateItems map[string]interface{}) (*model.Sala, error) {
	tx := r.db.Model(sala).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrSalaNaoEncontrada
	}

	return sala, nil
}

func (r *salaRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Sala{}, "id = ?", id).Error
}
