package repository

import (
	"cinema_digital_go/api/app/ingresso/model"
	salaModel "cinema_digital_go/api/app/sala/model"
	sessaoModel "cinema_digital_go/api/app/sessao/model"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sort"
)

type IngressoRepository interface {
	PoltronasOcupadas(idSessao uuid.UUID, poltronas []string) ([]model.Ingresso, error)
	Create(ingressos []model.Ingresso) error
	ListarPoltronasDisponiveis(sessao sessaoModel.Sessao, sala salaModel.Sala) ([]string, error)
}

type ingressoRepositoryImpl struct {
	db *gorm.DB
}

func NewIngressoRepository(db *gorm.DB) IngressoRepository {
	return &ingressoRepositoryImpl{db: db}
}

func (r *ingressoRepositoryImpl) PoltronasOcupadas(idSessao uuid.UUID, poltronas []string) ([]model.Ingresso, error) {
	var ocupadas []model.Ingresso
	err := r.db.Where("id_sessao = ? AND poltrona IN ?", idSessao, poltronas).Find(&ocupadas).Error
	return ocupadas, err
}

func (r *ingressoRepositoryImpl) Create(ingressos []model.Ingresso) error {
	return r.db.Create(&ingressos).Error
}

func (r *ingressoRepositoryImpl) ListarPoltronasDisponiveis(sessao sessaoModel.Sessao, sala salaModel.Sala) ([]string, error) {
	qtdFileiras := int(sala.Fileiras[0]-'A') + 1
	if qtdFileiras <= 0 {
		return nil, fmt.Errorf("fileiras inválidas")
	}
	poltronasPorFileira := sala.Poltronas / qtdFileiras

	validas := map[string]bool{}
	for i := 0; i < qtdFileiras; i++ {
		letra := string(rune('A' + i))
		for j := 1; j <= poltronasPorFileira; j++ {
			validas[fmt.Sprintf("%s%d", letra, j)] = true
		}
	}

	var ocupadas []model.Ingresso
	err := r.db.Where("id_sessao = ?", sessao.Id).Find(&ocupadas).Error
	if err != nil {
		return nil, err
	}

	for _, ing := range ocupadas {
		delete(validas, ing.Poltrona)
	}

	var disponiveis []string
	for p := range validas {
		disponiveis = append(disponiveis, p)
	}

	sort.Strings(disponiveis)
	return disponiveis, nil
}
