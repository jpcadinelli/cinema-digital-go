package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"unicode"
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

func GerarCamposAtualizacao(entidade any) map[string]interface{} {
	updateFields := make(map[string]interface{})
	val := reflect.ValueOf(entidade).Elem()

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface()
		fieldName := strings.ToLower(FormatarCamelCaseParaSnakeCase(val.Type().Field(i).Name))

		if !isZeroValue(fieldValue) {
			updateFields[fieldName] = fieldValue
		}
	}

	return updateFields
}

func isZeroValue(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func FormatarCamelCaseParaSnakeCase(str string) string {
	var resultado strings.Builder

	for i, char := range str {
		if unicode.IsUpper(char) && i != 0 {
			resultado.WriteRune('_')
		}
		resultado.WriteRune(unicode.ToLower(char))
	}

	return resultado.String()
}
