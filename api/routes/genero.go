package routes

import (
	dropResource "cinema_digital_go/api/app/dropdown/resource"
	"cinema_digital_go/api/app/genero/resource"
	"github.com/gin-gonic/gin"
)

func generoRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar)
	r.GET(routeDropdown, dropResource.DropdownGeneros)
}
