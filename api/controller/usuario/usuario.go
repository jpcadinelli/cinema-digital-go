package usuario

import (
	dbConetion "cinema_digital_go/api/database/conection"
	"cinema_digital_go/api/global/enum"
	"cinema_digital_go/api/global/erros"
	"cinema_digital_go/api/middleware"
	"cinema_digital_go/api/models"
	"cinema_digital_go/api/repository"
	"cinema_digital_go/api/security"
	"cinema_digital_go/api/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Criar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var u models.Usuario

	if err = ginctx.ShouldBindJSON(&u); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	u.Password = security.SHA256Encoder(u.Password)

	if err = repository.NewUsuarioRepository(dbConetion.DB).Create(&u); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusCreated, middleware.NewResponseBridge(nil, userResponse))
}

func Visualizar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioVisualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	u, err := repository.NewUsuarioRepository(dbConetion.DB).FindById(id, "Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, userResponse))
}

func Listar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioListar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	usuarios, err := repository.NewUsuarioRepository(dbConetion.DB).FindAll("Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	response := []*models.UsuarioDTOResponse{}
	for _, u := range usuarios {
		response = append(response, u.UsuarioToDTOResponse())
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, response))
}

func Dropdown(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDropdown) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	usuarios, err := repository.NewUsuarioRepository(dbConetion.DB).FindAll()
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	response := []*models.DropdownUUID{}
	for _, u := range usuarios {
		response = append(response, u.UsuarioToDropdownUUID())
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, response))
}

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var u models.Usuario

	if err = ginctx.ShouldBindJSON(&u); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if usuarioLogado.Id != u.Id {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(erros.ErrNaoPodeMudadarDadosDeOutroUsuario, nil))
		return
	}

	uOld, err := repository.NewUsuarioRepository(dbConetion.DB).FindById(u.Id)
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

	uOld, err = repository.NewUsuarioRepository(dbConetion.DB).Update(uOld, updateItems)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	reponse := uOld.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, reponse))
}

func Deletar(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDeletar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewUsuarioRepository(dbConetion.DB).Delete(id); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}

func AtribuirPermissao(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioAtribuirPermissao) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	idPermissaoStr := ginctx.Param("idPermissao")
	idPermissao, err := uuid.Parse(idPermissaoStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	permissaoUsuario, err := repository.NewPermissaoUsuarioRepository(dbConetion.DB).FindRelations(id, idPermissao)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, permissaoUsuario))
}

func RemoverPermissao(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioRemoverPermissao) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	idStr := ginctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	idPermissaoStr := ginctx.Param("idPermissao")
	idPermissao, err := uuid.Parse(idPermissaoStr)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	if err = repository.NewPermissaoUsuarioRepository(dbConetion.DB).Delete(id, idPermissao); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}

func VisualizarUsuarioLogado(ginctx *gin.Context) {
	usuarioLogado, err := service.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	u, err := repository.NewUsuarioRepository(dbConetion.DB).FindById(usuarioLogado.Id, "Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	userResponse := u.UsuarioToDTOResponse()
	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, userResponse))
}
