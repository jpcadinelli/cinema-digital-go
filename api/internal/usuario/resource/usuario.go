package resource

import (
	"cinema_digital_go/api/internal/dropdown/model"
	repository2 "cinema_digital_go/api/internal/permissao/resource"
	models2 "cinema_digital_go/api/internal/usuario/model"
	"cinema_digital_go/api/internal/usuario/repository"
	dbConetion "cinema_digital_go/api/pkg/database/conection"
	"cinema_digital_go/api/pkg/global/enum"
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/middleware"
	"cinema_digital_go/api/pkg/security"
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

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioCriar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var u models2.Usuario

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
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioVisualizar) {
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
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioListar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	usuarios, err := repository.NewUsuarioRepository(dbConetion.DB).FindAll("Permissoes")
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	response := []*models2.UsuarioDTOResponse{}
	for _, u := range usuarios {
		response = append(response, u.UsuarioToDTOResponse())
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, response))
}

func Dropdown(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDropdown) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	usuarios, err := repository.NewUsuarioRepository(dbConetion.DB).FindAll()
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

func Atualizar(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioAtualizar) {
		ginctx.JSON(http.StatusUnauthorized, middleware.NewResponseBridge(erros.ErrUsuarioNaoTemPermissao, nil))
		return
	}

	var u models2.Usuario

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
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioDeletar) {
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
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioAtribuirPermissao) {
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

	permissaoUsuario, err := repository2.NewPermissaoUsuarioRepository(dbConetion.DB).FindRelations(id, idPermissao)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, permissaoUsuario))
}

func RemoverPermissao(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	if !service2.VerificaPermissaoUsuario(*usuarioLogado, enum.PermissaoUsuarioRemoverPermissao) {
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

	if err = repository2.NewPermissaoUsuarioRepository(dbConetion.DB).Delete(id, idPermissao); err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, nil))
}

func VisualizarUsuarioLogado(ginctx *gin.Context) {
	usuarioLogado, err := service2.GetUsuarioLogado(ginctx)
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
