package resource

import (
	"cinema_digital_go/api/app/sessao/model"
	"cinema_digital_go/api/app/sessao/repository"
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

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoSessaoCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var s model.Sessao
	if err = ginctx.ShouldBindJSON(&s); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = s.Validar(); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewSessaoRepository(dbConection.DB).Create(&s); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, s))
}

func Visualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoSessaoAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	s, err := repository.NewSessaoRepository(dbConection.DB).FindById(*id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, s))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoSessaoAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	sOld, err := repository.NewSessaoRepository(dbConection.DB).FindById(*id)
	if err != nil {
		ginctx.JSON(http.StatusNotFound, middleware.NewResponseBridge(errors.New("Sessao não encontrado"), nil))
		return
	}

	var s model.Sessao
	if err = ginctx.ShouldBindJSON(&s); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := utils.GerarCamposAtualizacao(&s)

	sOld, err = repository.NewSessaoRepository(dbConection.DB).Update(sOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, sOld))
}

func Deletar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoSessaoDeletar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	err = repository.NewSessaoRepository(dbConection.DB).Delete(*id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusNoContent, middleware.NewResponseBridge(nil, nil))
}

func Listar(ginctx *gin.Context) {
	sessoes, err := repository.NewSessaoRepository(dbConection.DB).FindAll()
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, sessoes))
}
