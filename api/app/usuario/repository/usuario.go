package repository

import (
	"cinema_digital_go/api/app/permissao/model"
	models2 "cinema_digital_go/api/app/usuario/model"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*models2.Usuario, error)
	FindByEmail(email string) (*models2.Usuario, error)
	FindAll(preloads ...string) ([]models2.Usuario, error)
	Create(usuario *models2.Usuario) error
	Update(usuario *models2.Usuario, updateItems map[string]interface{}) (*models2.Usuario, error)
	Delete(id uuid.UUID) error
}

type usuarioRepositoryImpl struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepositoryImpl{db: db}
}

func (r *usuarioRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*models2.Usuario, error) {
	var usuario models2.Usuario

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

func (r *usuarioRepositoryImpl) FindByEmail(email string) (*models2.Usuario, error) {
	var usuario models2.Usuario

	tx := r.db.First(&usuario, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, erros.ErrUsuarioNaoEncontrado
	}

	return &usuario, nil
}

func (r *usuarioRepositoryImpl) FindAll(preloads ...string) ([]models2.Usuario, error) {
	var usuarios []models2.Usuario

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

func (r *usuarioRepositoryImpl) Create(usuario *models2.Usuario) error {
	return r.db.Create(usuario).Error
}

func (r *usuarioRepositoryImpl) Update(usuario *models2.Usuario, updateItems map[string]interface{}) (*models2.Usuario, error) {
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
		if err := tx.Where("id_usuario = ?", id).Delete(&model.PermissaoUsuario{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models2.Usuario{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}
