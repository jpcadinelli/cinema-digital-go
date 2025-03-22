package resource

import (
	"cinema_digital_go/api/app/filme/model"
	"cinema_digital_go/api/app/filme/repository"
	dbConection "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	"cinema_digital_go/api/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoFilmeCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var f model.Filme
	f.IdUsuarioRegistro = usuarioLogado.Id

	if err = ginctx.ShouldBindJSON(&f); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewFilmeRepository(dbConection.DB).Create(&f); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, f))
}

func Visualizar(ginctx *gin.Context) {
	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	f, err := repository.NewFilmeRepository(dbConection.DB).FindById(*id, "Generos")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, f))
}

func Listar(ginctx *gin.Context) {
	filmes, err := repository.NewFilmeRepository(dbConection.DB).FindAll("Generos")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, filmes))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoFilmeAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	fOld, err := repository.NewFilmeRepository(dbConection.DB).FindById(*id)
	if err != nil {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(errors.New("Filme não encontrado"), nil))
		return
	}

	var f model.Filme
	if err = ginctx.ShouldBindJSON(&f); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := utils.GerarCamposAtualizacao(&f)

	fOld, err = repository.NewFilmeRepository(dbConection.DB).Update(fOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, fOld))
}

func Deletar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoFilmeDeletar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	err = repository.NewFilmeRepository(dbConection.DB).Delete(*id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusNoContent, middleware.NewResponseBridge(nil, nil))
}
