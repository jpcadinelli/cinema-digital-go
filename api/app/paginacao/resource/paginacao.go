package resource

import (
	"cinema_digital_go/api/app/paginacao/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

func PaginarConsulta[P any](c *gin.Context, query *gorm.DB, out *[]P) (model.Paginacao, error) {
	pagina, limite, offset, err := getPaginationParams(c)
	if err != nil {
		return model.Paginacao{}, err
	}

	var totalItens int64
	if err = query.Count(&totalItens).Error; err != nil {
		return model.Paginacao{}, err
	}

	if err = query.Offset(offset).Limit(limite).Find(out).Error; err != nil {
		return model.Paginacao{}, err
	}

	totalPaginas := int(math.Ceil(float64(totalItens) / float64(limite)))

	meta := model.Paginacao{
		TotalItens:   int(totalItens),
		Limite:       limite,
		PaginaAtual:  pagina,
		TotalPaginas: totalPaginas,
	}

	return meta, nil
}

func getPaginationParams(c *gin.Context) (int, int, int, error) {
	paginaStr := c.DefaultQuery("pagina", "1")
	limiteStr := c.DefaultQuery("limite", "10")

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
