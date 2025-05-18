package resource

import (
	"cinema_digital_go/api/app/ingresso/model"
	"cinema_digital_go/api/app/ingresso/repository"
	salaModel "cinema_digital_go/api/app/sala/model"
	sessaoModel "cinema_digital_go/api/app/sessao/model"
	dbConection "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/middleware"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ComprarIngresso(ginctx *gin.Context) {
	var req model.CompraIngressoRequest
	if err := ginctx.ShouldBindJSON(&req); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	db := dbConection.DB
	repo := repository.NewIngressoRepository(db)

	var sessao sessaoModel.Sessao
	if err := db.First(&sessao, "id = ?", req.IdSessao).Error; err != nil {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(fmt.Errorf("sessão não encontrada"), nil))
		return
	}
	if sessao.Disponibilidade != 1 {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(fmt.Errorf("sessão indisponível"), nil))
		return
	}

	var sala salaModel.Sala
	if err := db.First(&sala, "id = ?", sessao.IdSala).Error; err != nil {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(fmt.Errorf("sala não encontrada"), nil))
		return
	}

	qtdFileiras := int(sala.Fileiras[0]-'A') + 1
	if qtdFileiras <= 0 {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(fmt.Errorf("fileiras inválidas"), nil))
		return
	}
	poltronasPorFileira := sala.Poltronas / qtdFileiras

	validas := map[string]bool{}
	for i := 0; i < qtdFileiras; i++ {
		letra := string(rune('A' + i))
		for j := 1; j <= poltronasPorFileira; j++ {
			validas[fmt.Sprintf("%s%d", letra, j)] = true
		}
	}

	for _, p := range req.Poltronas {
		if !validas[p] {
			ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(fmt.Errorf("poltrona inválida: %s", p), nil))
			return
		}
	}

	ocupadas, err := repo.PoltronasOcupadas(req.IdSessao, req.Poltronas)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}
	if len(ocupadas) > 0 {
		lista := []string{}
		for _, o := range ocupadas {
			lista = append(lista, o.Poltrona)
		}
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(fmt.Errorf("poltronas já ocupadas: %s", strings.Join(lista, ", ")), nil))
		return
	}

	var ingressos []model.Ingresso
	for _, p := range req.Poltronas {
		ingressos = append(ingressos, model.Ingresso{
			Id:         uuid.New(),
			IdSessao:   req.IdSessao,
			IdUsuario:  req.IdUsuario,
			Poltrona:   p,
			CompradoEm: time.Now(),
		})
	}
	if err = repo.Create(ingressos); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, ingressos))
}

func ListarPoltronasDisponiveis(ginctx *gin.Context) {
	idSessaoParam := ginctx.Param("idSessao")
	idSessao, err := uuid.Parse(idSessaoParam)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(fmt.Errorf("ID de sessão inválido"), nil))
		return
	}

	db := dbConection.DB
	repo := repository.NewIngressoRepository(db)

	var sessao sessaoModel.Sessao
	if err := db.First(&sessao, "id = ?", idSessao).Error; err != nil {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(fmt.Errorf("sessão não encontrada"), nil))
		return
	}

	var sala salaModel.Sala
	if err := db.First(&sala, "id = ?", sessao.IdSala).Error; err != nil {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(fmt.Errorf("sala não encontrada"), nil))
		return
	}

	disponiveis, err := repo.ListarPoltronasDisponiveis(sessao, sala)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, disponiveis))
}
