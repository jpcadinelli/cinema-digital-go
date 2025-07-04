package routes

import (
	"cinema_digital_go/api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

const (
	route         = ""
	routeId       = "/:id"
	routeFiltro   = "/filtro"
	routeDropdown = "/dropdown"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		usuarioGroup := main.Group("/usuarios")
		{
			usuarioRoutes(usuarioGroup)
		}
		loginGroup := main.Group("/login")
		{
			loginRoutes(loginGroup)
		}
		permissaoGroup := main.Group("/permissoes", middleware.Auth())
		{
			permissaoRoutes(permissaoGroup)
		}
		filmeGroup := main.Group("/filmes")
		{
			filmeRoutes(filmeGroup)
		}
		generoGroup := main.Group("/generos", middleware.Auth())
		{
			generoRoutes(generoGroup)
		}
		salaGroup := main.Group("/salas", middleware.Auth())
		{
			salaRoutes(salaGroup)
		}
		sessaoGroup := main.Group("/sessoes")
		{
			sessaoRoutes(sessaoGroup)
		}
		ingressoGroup := main.Group("/ingressos", middleware.Auth())
		{
			ingressoRoutes(ingressoGroup)
		}
	}

	return router
}
