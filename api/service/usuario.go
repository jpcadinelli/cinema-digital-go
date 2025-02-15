package service

import (
	dbConetion "cinema_digital_go/api/database/conection"
	"cinema_digital_go/api/global/erros"
	"cinema_digital_go/api/models"
	"cinema_digital_go/api/repository"
	"cinema_digital_go/api/security"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	BearerSchema = "Bearer "
)

func GetIdUsuarioLogado(ginctx *gin.Context) (uuid.UUID, error) {
	var (
		id  uuid.UUID
		err error
	)

	header := ginctx.Request.Header.Get("Authorization")
	if header == "" {
		return id, erros.ErrTokenInexistente
	}

	token := header[len(BearerSchema):]

	if id, err = security.NewJWTService().GetUserId(token); err != nil {
		return id, err
	}

	return id, nil
}

func GetUsuarioLogado(ginctx *gin.Context) (*models.UsuarioDTOResponse, error) {
	header := ginctx.Request.Header.Get("Authorization")
	if header == "" {
		return nil, erros.ErrTokenInexistente
	}

	token := header[len(BearerSchema):]

	id, err := security.NewJWTService().GetUserId(token)
	if err != nil {
		return nil, err
	}

	usuario, err := repository.NewUsuarioRepository(dbConetion.DB).FindById(id, "Permissoes")
	if err != nil {
		return nil, err
	}

	userResponse := usuario.UsuarioToDTOResponse()
	return userResponse, nil
}
