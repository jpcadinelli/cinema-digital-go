package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetParamID(params gin.Params, name string) (paramID *uuid.UUID, err error) {
	if id, ok := params.Get(name); ok {
		var parsedID uuid.UUID
		parsedID, err = uuid.Parse(id)
		if err != nil {
			return nil, errors.New("formato de UUID inválido")
		}
		paramID = &parsedID
		return paramID, nil
	}
	return nil, errors.New("parâmetro 'id' não encontrado")
}
