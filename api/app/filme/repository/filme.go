package repository

import (
	"cinema_digital_go/api/app/filme/model"
	modelGen "cinema_digital_go/api/app/genero/model"
	modelPag "cinema_digital_go/api/app/paginacao/model"
	"cinema_digital_go/api/app/paginacao/repository"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmeRepository interface {
	FindById(id uuid.UUID, preloads ...string) (*model.Filme, error)
	FindAll(preloads ...string) ([]model.Filme, error)
	Create(filme *model.Filme) error
	Update(filme *model.Filme) (*model.Filme, error)
	Delete(id uuid.UUID) error
	List(ginctx *gin.Context) (*modelPag.Paginacao, error)
	GetEmCartaz() ([]model.Filme, error)
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
	return r.db.Transaction(func(tx *gorm.DB) error {
		f := filme.GetOnlyFilme()
		if err := tx.Create(f).Error; err != nil {
			return err
		}

		filme.Id = f.Id

		var listRe []model.ReFilmeGenero
		for _, re := range filme.Generos {
			listRe = append(listRe, model.ReFilmeGenero{
				IdFilme:  f.Id,
				IdGenero: re.Id,
			})
		}
		if err := tx.Create(&listRe).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *filmeRepositoryImpl) Update(filme *model.Filme) (*model.Filme, error) {
	var (
		err       error
		listReOld []model.ReFilmeGenero
		listReAdd []model.ReFilmeGenero
		listReRem []model.ReFilmeGenero
	)

	if err = r.db.Model(model.ReFilmeGenero{}).Where("id_filme = ?", filme.Id).Find(&listReOld).Error; err != nil {
		return nil, err
	}

	oldReMap := make(map[uuid.UUID]model.ReFilmeGenero)
	for _, re := range listReOld {
		oldReMap[re.IdGenero] = re
	}

	for _, genero := range filme.Generos {
		if _, exists := oldReMap[genero.Id]; !exists {
			listReAdd = append(listReAdd, model.ReFilmeGenero{
				Id:       uuid.New(),
				IdFilme:  filme.Id,
				IdGenero: genero.Id,
			})
		}
	}

	newGenMap := make(map[uuid.UUID]modelGen.Genero)
	for _, genero := range filme.Generos {
		newGenMap[genero.Id] = genero
	}

	for _, re := range listReOld {
		if _, exists := newGenMap[re.IdGenero]; !exists {
			listReRem = append(listReRem, re)
		}
	}

	updateItems := map[string]interface{}{
		"titulo":         filme.Titulo,
		"sinopse":        filme.Sinopse,
		"diretor":        filme.Diretor,
		"duracao":        filme.Duracao,
		"ano_lancamento": filme.AnoLancamento,
		"classificacao":  filme.Classificacao,
		"nota":           filme.Nota,
	}

	txDB := r.db.First(filme, "id = ?", filme.Id)
	if txDB.Error != nil {
		return nil, txDB.Error
	}
	if txDB.RowsAffected == 0 {
		return nil, erros.ErrFilmeNaoEncontrado
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {

		if err = tx.Model(&model.Filme{}).Where("id = ?", filme.Id).Updates(updateItems).Error; err != nil {
			return err
		}

		if len(listReAdd) > 0 {
			if err = tx.Create(&listReAdd).Error; err != nil {
				return err
			}
		}

		for _, re := range listReRem {
			if err = tx.Delete(&model.ReFilmeGenero{}, "id = ?", re.Id).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return filme, err
}

func (r *filmeRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Filme{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *filmeRepositoryImpl) List(ginctx *gin.Context) (*modelPag.Paginacao, error) {
	query := r.db.Model(&model.Filme{}).Preload("Generos")

	var filmes []model.Filme
	return repository.ConsultaPaginada(ginctx, query, &filmes)
}

func (r *filmeRepositoryImpl) GetEmCartaz() ([]model.Filme, error) {
	var filmes []model.Filme

	tx := r.db.
		Preload("Generos").
		Joins("JOIN sessao ON sessao.id_filme = filme.id").
		Where("sessao.disponibilidade = ?", enum.SessaoDisponivel).
		Find(&filmes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return filmes, erros.ErrFilmeNaoEncontrado
	}

	return filmes, nil
}
