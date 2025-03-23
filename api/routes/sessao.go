package routes

import (
	"cinema_digital_go/api/app/sessao/resource"
	"github.com/gin-gonic/gin"
)

func sessaoRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
}
