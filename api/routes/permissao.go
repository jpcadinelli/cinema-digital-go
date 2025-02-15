package routes

import (
	"cinema_digital_go/api/controller/permissao"
	"github.com/gin-gonic/gin"
)

func permissaoRoutes(r *gin.RouterGroup) {
	r.POST(route, permissao.Criar)
	r.GET(routeId, permissao.Visualizar)
	r.GET(route, permissao.Listar)
	r.GET(routeDropdown, permissao.Dropdown)
	r.PUT(route, permissao.Atualizar)
	r.DELETE(routeId, permissao.Deletar)
}
