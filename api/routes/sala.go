package routes

import (
	dropResource "cinema_digital_go/api/app/dropdown/resource"
	"cinema_digital_go/api/app/sala/resource"
	"github.com/gin-gonic/gin"
)

func salaRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar)
	r.GET(routeDropdown, dropResource.DropdownSalas)
	r.PUT(routeId, resource.Atualizar)
	r.DELETE(routeId, resource.Deletar)
}
