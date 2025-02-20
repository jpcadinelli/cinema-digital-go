package routes

import (
	"cinema_digital_go/api/app/usuario/resource"
	"github.com/gin-gonic/gin"
)

func usuarioRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
	r.GET(routeId, resource.Visualizar)
	r.GET(route, resource.Listar)
	r.GET(routeDropdown, resource.Dropdown)
	r.PUT(route, resource.Atualizar)
	r.DELETE(routeId, resource.Deletar)

	r.POST(routeId+"/permissao/:idPermissao", resource.AtribuirPermissao)
	r.DELETE(routeId+"/permissao/:idPermissao", resource.RemoverPermissao)

	r.GET(route+"logado", resource.VisualizarUsuarioLogado)
}
