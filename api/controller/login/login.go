package login

import (
	dbConection "cinema_digital_go/api/database/conection"
	"cinema_digital_go/api/global/erros"
	"cinema_digital_go/api/middleware"
	"cinema_digital_go/api/models"
	"cinema_digital_go/api/repository"
	"cinema_digital_go/api/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ginctx *gin.Context) {
	var l models.Login

	if err := ginctx.ShouldBindJSON(&l); err != nil {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(err, nil))
		return
	}

	usuario, err := repository.NewUsuarioRepository(dbConection.DB).FindByEmail(l.Email)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	if usuario.Password != security.SHA256Encoder(l.Password) {
		ginctx.JSON(http.StatusBadRequest, middleware.NewResponseBridge(erros.ErrCredenciaisInvalidas, nil))
		return
	}

	token, err := security.NewJWTService().GenerateToken(usuario.Id)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, middleware.NewResponseBridge(err, nil))
		return
	}

	ginctx.JSON(http.StatusOK, middleware.NewResponseBridge(nil, token))
}
