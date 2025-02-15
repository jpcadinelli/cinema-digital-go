package routes

import (
	"cinema_digital_go/api/controller/login"
	"github.com/gin-gonic/gin"
)

func loginRoutes(r *gin.RouterGroup) {
	r.POST(route, login.Login)
}
