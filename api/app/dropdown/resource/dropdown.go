package resource

import (
	"cinema_digital_go/api/app/dropdown/model"
	permissaoRepository "cinema_digital_go/api/app/permissao/repository"
	usuarioRepository "cinema_digital_go/api/app/usuario/repository"
	dbConection "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	"cinema_digital_go/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DropdownPermissao(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoPermissaoDropdown) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	permissoes, err := permissaoRepository.NewPermissaoRepository(dbConection.DB).FindAll()
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

func DropdownUsuarios(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDropdown) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	usuarios, err := usuarioRepository.NewUsuarioRepository(dbConection.DB).FindAll()
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	response := []*model.DropdownUUID{}
	for _, u := range usuarios {
		response = append(response, u.UsuarioToDropdownUUID())
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, response))
}
