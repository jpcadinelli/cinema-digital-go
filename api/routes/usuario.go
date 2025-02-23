package routes

import (
	"cinema_digital_go/api/app/usuario/resource"
	"cinema_digital_go/api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func usuarioRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar, middleware.Auth(), middleware.Auth())
	r.GET(route, resource.Listar, middleware.Auth())
	r.GET(routeDropdown, resource.Dropdown, middleware.Auth())
	r.PUT(route, resource.Atualizar, middleware.Auth())
	r.DELETE(routeId, resource.Deletar, middleware.Auth())

	r.POST(routeId+"/permissao/:idPermissao", resource.AtribuirPermissao, middleware.Auth())
	r.DELETE(routeId+"/permissao/:idPermissao", resource.RemoverPermissao, middleware.Auth())

	r.GET(route+"logado", resource.VisualizarUsuarioLogado)
}
