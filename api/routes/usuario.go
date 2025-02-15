package routes

import (
	"cinema_digital_go/api/controller/usuario"
	"github.com/gin-gonic/gin"
)

func usuarioRoutes(r *gin.RouterGroup) {
	r.POST(route, usuario.Criar)
	r.GET(routeId, usuario.Visualizar)
	r.GET(route, usuario.Listar)
	r.GET(routeDropdown, usuario.Dropdown)
	r.PUT(route, usuario.Atualizar)
	r.DELETE(routeId, usuario.Deletar)

	r.POST(routeId+"/permissao/:idPermissao", usuario.AtribuirPermissao)
	r.DELETE(routeId+"/permissao/:idPermissao", usuario.RemoverPermissao)

	r.GET(route+"logado", usuario.VisualizarUsuarioLogado)
}
