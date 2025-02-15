package repository

import (
	"cinema_digital_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissaoUsuarioRepository interface {
	Create(permissaoUsuario *models.PermissaoUsuario) error
	FindRelations(id, idPermissao uuid.UUID) (*models.PermissaoUsuario, error)
	Delete(id, idPermissao uuid.UUID) error
}

type permissaoUsuarioRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissaoUsuarioRepository(db *gorm.DB) PermissaoUsuarioRepository {
	return &permissaoUsuarioRepositoryImpl{db: db}
}

func (r *permissaoUsuarioRepositoryImpl) Create(permissaoUsuario *models.PermissaoUsuario) error {
	return r.db.Create(permissaoUsuario).Error
}

func (r *permissaoUsuarioRepositoryImpl) FindRelations(id, idPermissao uuid.UUID) (*models.PermissaoUsuario, error) {
	var permissaoUsuario models.PermissaoUsuario

	tx := r.db.Find(&permissaoUsuario, "id_usuario = ? AND id_permissao = ?", id, idPermissao)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		permissaoUsuario.IdUsuario = id
		permissaoUsuario.IdPermissao = idPermissao

		if err := r.Create(&permissaoUsuario); err != nil {
			return nil, err
		}
	}

	return &permissaoUsuario, nil
}

func (r *permissaoUsuarioRepositoryImpl) Delete(id, idPermissao uuid.UUID) error {
	tx := r.db.Delete(&models.PermissaoUsuario{}, "id_usuario = ? AND id_permissao = ?", id, idPermissao)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
