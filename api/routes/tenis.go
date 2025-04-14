package routes

import (
	"cinema_digital_go/api/app/tenis/resource"
	"github.com/gin-gonic/gin"
)

func tenisRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Criar)
}
