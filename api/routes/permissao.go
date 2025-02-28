package routes

import (
	dropdown "cinema_digital_go/api/app/dropdown/resource"
	"cinema_digital_go/api/app/permissao/resource"
	"github.com/gin-gonic/gin"
)

func permissaoRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar)
	r.GET(route, resource.Listar)
	r.GET(routeDropdown, dropdown.DropdownPermissao)
	r.PUT(route, resource.Atualizar)
	r.DELETE(routeId, resource.Deletar)
}
