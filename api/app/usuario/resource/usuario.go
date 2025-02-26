package resource

import (
	permissaoRepository "cinema_digital_go/api/app/permissao/repository"
	permissao "cinema_digital_go/api/app/permissao/resource"
	"cinema_digital_go/api/app/usuario/model"
	"cinema_digital_go/api/app/usuario/repository"
	dbConection "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	"cinema_digital_go/api/pkg/security"
	"cinema_digital_go/api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	var (
		u   model.Usuario
		err error
	)

	if err = ginctx.ShouldBindJSON(&u); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	u.Password = security.SHA256Encoder(u.Password)
	if u.Permissoes, err = permissaoRepository.NewPermissaoRepository(dbConection.DB).FindByGroup(enum.GrupoN1Permissoes); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewUsuarioRepository(dbConection.DB).Create(&u); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, userResponse))
}

func Visualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioVisualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	u, err := repository.NewUsuarioRepository(dbConection.DB).FindById(*id, "Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, userResponse))
}

func Listar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioListar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	usuarios, err := repository.NewUsuarioRepository(dbConection.DB).FindAll("Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	var response []*model.UsuarioDTOResponse
	for _, u := range usuarios {
		response = append(response, u.UsuarioToDTOResponse())
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, response))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var u model.Usuario

	if err = ginctx.ShouldBindJSON(&u); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if usuarioLogado.Id != u.Id {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(erros.ErrNaoPodeMudadarDadosDeOutroUsuario, nil))
		return
	}

	uOld, err := repository.NewUsuarioRepository(dbConection.DB).FindById(u.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	updateItems := map[string]interface{}{
		"primeiro_nome": u.PrimeiroNome,
		"ultimo_nome":   u.UltimoNome,
		"email":         u.Email,
		"password":      security.SHA256Encoder(u.Password),
	}

	uOld, err = repository.NewUsuarioRepository(dbConection.DB).Update(uOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	reponse := uOld.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, reponse))
}

func Deletar(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDeletar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	if err = repository.NewUsuarioRepository(dbConection.DB).Delete(*id); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}

func AtribuirPermissao(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioAtribuirPermissao) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	idPermissao, err := utils.GetParamID(ginctx.Params, "idPermissao")
	if err != nil {
		return
	}

	permissaoUsuario, err := permissao.NewPermissaoUsuarioRepository(dbConection.DB).FindRelations(*id, *idPermissao)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, permissaoUsuario))
}

func RemoverPermissao(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !utils.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioRemoverPermissao) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	id, err := utils.GetParamID(ginctx.Params, "id")
	if err != nil {
		return
	}

	idPermissao, err := utils.GetParamID(ginctx.Params, "idPermissao")
	if err != nil {
		return
	}

	if err = permissao.NewPermissaoUsuarioRepository(dbConection.DB).Delete(*id, *idPermissao); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}

func VisualizarUsuarioLogado(ginctx *gin.Context) {
	usuarioLogado, err := utils.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	u, err := repository.NewUsuarioRepository(dbConection.DB).FindById(usuarioLogado.Id, "Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, userResponse))
}
