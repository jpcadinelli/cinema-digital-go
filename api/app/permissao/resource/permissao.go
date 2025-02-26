package resource

import (
	"cinema_digital_go/api/app/permissao/model"
	"cinema_digital_go/api/app/permissao/repository"
	dbConection "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	"cinema_digital_go/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var p model.Permissao

	if err = ginctx.ShouldBindJSON(&p); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewPermissaoRepository(dbConection.DB).Create(&p); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, p))
}

func Visualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoVisualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	p, err := repository.NewPermissaoRepository(dbConection.DB).FindById(*id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, p))
}

func Listar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoListar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	permissoes, err := repository.NewPermissaoRepository(dbConection.DB).FindAll()
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, permissoes))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var p model.Permissao

	if err = ginctx.ShouldBindJSON(&p); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	pOld, err := repository.NewPermissaoRepository(dbConection.DB).FindById(p.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := map[string]interface{}{
		"nome":      p.Nome,
		"descricao": p.Descricao,
	}

	pOld, err = repository.NewPermissaoRepository(dbConection.DB).Update(pOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, pOld))
}

func Deletar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoDeletar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	err = repository.NewPermissaoRepository(dbConection.DB).Delete(*id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}
