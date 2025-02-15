package repository

import (
	"cinema_digital_go/api/global/erros"
	"cinema_digital_go/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models.Usuario, error)
	FindByEmail(email string) (*models.Usuario, error)
	FindAll(preloads ...string) ([]models.Usuario, error)
	Create(usuario *models.Usuario) error
	Update(usuario *models.Usuario, updateItems map[string]interface{}) (*models.Usuario, error)
	Delete(id uuid.UUID) error
}

type usuarioRepositoryImpl struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepositoryImpl{db: db}
}

func (r *usuarioRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models.Usuario, error) {
	var usuario models.Usuario

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.First(&usuario, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrUsuarioNaoEncontrado
	}

	return &usuario, nil
}

func (r *usuarioRepositoryImpl) FindByEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario

	tx := r.db.First(&usuario, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrUsuarioNaoEncontrado
	}

	return &usuario, nil
}

func (r *usuarioRepositoryImpl) FindAll(preloads ...string) ([]models.Usuario, error) {
	var usuarios []models.Usuario

	tx := r.db
	if len(preloads) > 0 {
		for _, preload := range preloads {
			tx = tx.Preload(preload)
		}
	}

	tx = tx.Find(&usuarios)
	if tx.Error != nil {
		return usuarios, tx.Error
	}
	if tx.RowsAffected == 0 {
		return usuarios, erros.ErrUsuarioNaoEncontrado
	}

	return usuarios, nil
}

func (r *usuarioRepositoryImpl) Create(usuario *models.Usuario) error {
	return r.db.Create(usuario).Error
}

func (r *usuarioRepositoryImpl) Update(usuario *models.Usuario, updateItems map[string]interface{}) (*models.Usuario, error) {
	tx := r.db.Model(usuario).Updates(updateItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrUsuarioNaoEncontrado
	}

	return usuario, nil
}

func (r *usuarioRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id_usuario = ?", id).Delete(&models.PermissaoUsuario{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Usuario{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}
