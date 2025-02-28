package routes

import (
	"cinema_digital_go/api/app/filme/resource"
	"github.com/gin-gonic/gin"
)

func filmeRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar)
	r.GET(route, resource.Listar)
	r.PUT(routeId, resource.Atualizar)
	r.DELETE(routeId, resource.Deletar)
}
