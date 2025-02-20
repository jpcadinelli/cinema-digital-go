package resource

import (
	"cinema_digital_go/api/app/dropdown/model"
	models2 "cinema_digital_go/api/app/permissao/model"
	"cinema_digital_go/api/app/permissao/repository"
	dbConetion "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	service2 "cinema_digital_go/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var p models2.Permissao

	if err = ginctx.ShouldBindJSON(&p); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewPermissaoRepository(dbConetion.DB).Create(&p); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, p))
}

func Visualizar(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoVisualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	p, err := repository.NewPermissaoRepository(dbConetion.DB).FindById(id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, p))
}

func Listar(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoListar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	permissoes, err := repository.NewPermissaoRepository(dbConetion.DB).FindAll()
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, permissoes))
}

func Dropdown(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoDropdown) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	permissoes, err := repository.NewPermissaoRepository(dbConetion.DB).FindAll()
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	var response []*model.DropdownUUID
	for _, p := range permissoes {
		response = append(response, p.PermissaoToDropdownUUID())
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, response))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var p models2.Permissao

	if err = ginctx.ShouldBindJSON(&p); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	pOld, err := repository.NewPermissaoRepository(dbConetion.DB).FindById(p.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := map[string]interface{}{
		"nome":      p.Nome,
		"descricao": p.Descricao,
	}

	pOld, err = repository.NewPermissaoRepository(dbConetion.DB).Update(pOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, pOld))
}

func Deletar(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoDeletar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	err = repository.NewPermissaoRepository(dbConetion.DB).Delete(id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}
