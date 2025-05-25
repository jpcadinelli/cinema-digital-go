package routes

import (
	"cinema_digital_go/api/app/ingresso/resource"
	"github.com/gin-gonic/gin"
)

func ingressoRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.ComprarIngresso)
	r.GET("/poltronas-disponiveis/:idSessao", resource.ListarPoltronasDisponiveis)
}
