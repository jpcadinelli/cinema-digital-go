package repository

import (
	"cinema_digital_go/api/app/permissao/model"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissaoRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*model.Permissao, error)
	FindAll(preloads ...string) ([]model.Permissao, error)
	FindByGroup(group []string) ([]model.Permissao, error)
	Create(permissao *model.Permissao) error
	Update(permissao *model.Permissao, updateItems map[string]interface{}) (*model.Permissao, error)
	Delete(id uuid.UUID) error
	GerenciaPermissoes() error
}

type permissaoRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissaoRepository(db *gorm.DB) PermissaoRepository {
	return &permissaoRepositoryImpl{db: db}
}

func (r *permissaoRepositoryImpl) FindById(id uuid.UUID, preloads ...string) (*model.Permissao, error) {
	var permissao model.Permissao

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

func (r *permissaoRepositoryImpl) FindAll(preloads ...string) ([]model.Permissao, error) {
	var permissoes []model.Permissao

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

func (r *permissaoRepositoryImpl) FindByGroup(group []string) ([]model.Permissao, error) {
	var permissoes []model.Permissao

	tx := r.db.Where("permissao.nome IN ?", group).Find(&permissoes)
	if tx.Error != nil {
		return permissoes, tx.Error
	}
	if tx.RowsAffected == 0 {
		return permissoes, erros.ErrGrupoDePermissoesNaoEncontradas
	}

	return permissoes, nil
}

func (r *permissaoRepositoryImpl) Create(permissao *model.Permissao) error {
	return r.db.Create(permissao).Error
}

func (r *permissaoRepositoryImpl) Update(permissao *model.Permissao, updateItems map[string]interface{}) (*model.Permissao, error) {
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
		if err := tx.Where("id_permissao = ?", id).Delete(&model.PermissaoUsuario{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.Permissao{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *permissaoRepositoryImpl) GerenciaPermissoes() error {
	var permissoes []model.Permissao

	tx := r.db.Find(&permissoes)
	if tx.Error != nil {
		return tx.Error
	}

	var permissoesFaltantes []model.Permissao
	for _, p := range enum.ListaPermissoes {
		faltante := true
		for _, permissao := range permissoes {
			if permissao.Nome == p {
				faltante = false
				break
			}
		}
		if faltante {
			permissoesFaltantes = append(permissoesFaltantes, model.Permissao{
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
