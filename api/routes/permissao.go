package routes

import (
	"cinema_digital_go/api/internal/permissao/resource"
	"github.com/gin-gonic/gin"
)

func permissaoRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar)
	r.GET(route, resource.Listar)
	r.GET(routeDropdown, resource.Dropdown)
	r.PUT(route, resource.Atualizar)
	r.DELETE(routeId, resource.Deletar)
}
