package repository

import (
	"cinema_digital_go/api/global/enum"
	"cinema_digital_go/api/global/erros"
	"cinema_digital_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissaoRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Permissao, error)
	FindAll(preloads ...string) ([]models.Permissao, error)
	Create(permissao *models.Permissao) error
	Update(permissao *models.Permissao, updateItems map[string]interface{}) (*models.Permissao, error)
	Delete(id uuid.UUID) error
	GerenciaPermissoes() error
}

type permissaoRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissaoRepository(db *gorm.DB) PermissaoRepository {
	return &permissaoRepositoryImpl{db: db}
}

func (r *permissaoRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Permissao, error) {
	var permissao models.Permissao

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&permissao, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrPermissaoNaoEncontrada
	}

	return &permissao, nil
}

func (r *permissaoRepositoryImpl) FindAll(preloads ...string) ([]models.Permissao, error) {
	var permissoes []models.Permissao

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&permissoes)
	if tx.Error != nil {
		return permissoes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return permissoes, erros.ErrPermissaoNaoEncontrada
	}

	return permissoes, nil
}

func (r *permissaoRepositoryImpl) Create(permissao *models.Permissao) error {
	return r.db.Create(permissao).Error
}

func (r *permissaoRepositoryImpl) Update(permissao *models.Permissao, updateItems map[string]interface{}) (*models.Permissao, error) {
	tx := r.db.Model(permissao).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrPermissaoNaoEncontrada
	}

	return permissao, nil
}

func (r *permissaoRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id_permissao = ?", id).Delete(&models.PermissaoUsuario{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Permissao{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *permissaoRepositoryImpl) GerenciaPermissoes() error {
	var permissoes []models.Permissao

	tx := r.db.Find(&permissoes)
	if tx.Error != nil {
		return tx.Error
	}

	var permissoesFaltantes []models.Permissao
	for _, p := range enum.ListaPermissoes {
		faltante := true
		for _, permissao := range permissoes {
			if permissao.Nome == p {
				faltante = false
				break
			}
		}
		if faltante {
			permissoesFaltantes = append(permissoesFaltantes, models.Permissao{
				Nome:      p,
				Descricao: "Criado pelo sistema.",
			})
		}
	}

	if len(permissoesFaltantes) > 0 {
		tx = tx.Create(&permissoesFaltantes)
	}

	return tx.Error
}
