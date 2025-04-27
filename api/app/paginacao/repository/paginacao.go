package repository

import (
	"cinema_digital_go/api/app/paginacao/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

func ConsultaPaginada[P any](ginctx *gin.Context, query *gorm.DB, baseModel *[]P) (*model.Paginacao, error) {
	pagina, limite, offset, err := getPaginationParams(ginctx)
	if err != nil {
		return nil, err
	}

	var totalItens int64
	if err = query.Count(&totalItens).Error; err != nil {
		return nil, err
	}

	if err = query.Offset(offset).Limit(limite).Find(baseModel).Error; err != nil {
		return nil, err
	}

	totalPaginas := int(math.Ceil(float64(totalItens) / float64(limite)))

	paginacao := &model.Paginacao{
		Conteudo:     baseModel,
		TotalItens:   int(totalItens),
		LimiteItens:  limite,
		PaginaAtual:  pagina,
		TotalPaginas: totalPaginas,
	}

	return paginacao, nil
}

func getPaginationParams(ginctx *gin.Context) (int, int, int, error) {
	paginaStr := ginctx.DefaultQuery("pagina", "1")
	limiteStr := ginctx.DefaultQuery("limite", "10")

	pagina, err := strconv.Atoi(paginaStr)
	if err != nil || pagina < 1 {
		return 0, 0, 0, fmt.Errorf("parâmetro 'pagina' inválido")
	}

	limite, err := strconv.Atoi(limiteStr)
	if err != nil || limite < 1 {
		return 0, 0, 0, fmt.Errorf("parâmetro 'limite' inválido")
	}

	offset := (pagina - 1) * limite
	return pagina, limite, offset, nil
}
