package routes

import (
	"cinema_digital_go/api/app/filme/resource"
	"cinema_digital_go/api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func filmeRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar, middleware.Auth())
	r.GET(routeId, resource.Visualizar, middleware.Auth())
	r.GET(route, resource.Listar, middleware.Auth())
	r.PUT(routeId, resource.Atualizar, middleware.Auth())
	r.DELETE(routeId, resource.Deletar, middleware.Auth())

	r.GET("em-cartaz", resource.EmCartaz)
}
