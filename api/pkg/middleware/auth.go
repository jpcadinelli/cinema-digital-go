package middleware

import (
	"cinema_digital_go/api/pkg/global/erros"
	"cinema_digital_go/api/pkg/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		const BearerSchema = "Bearer "
		header := ginctx.Request.Header.Get("Authorization")
		if header == "" {
			ginctx.AbortWithStatusJSON(http.StatusUnauthorized, NewResponseBridge(erros.ErrTokenInexistente, nil))
		}

		token := header[len(BearerSchema):]

		if !security.NewJWTService().ValidateToken(token) {
			ginctx.AbortWithStatusJSON(http.StatusUnauthorized, NewResponseBridge(erros.ErrTokenInvalido, nil))
		}
	}
}
