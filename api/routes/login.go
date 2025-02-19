package routes

import (
	"cinema_digital_go/api/internal/login/resource"
	"github.com/gin-gonic/gin"
)

func loginRoutes(r *gin.RouterGroup) {
	r.POST(route, resource.Login)
}
