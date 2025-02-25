package routes

import (
	"cinema_digital_go/api/app/genero/repository"
	"cinema_digital_go/api/app/genero/resource"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGeneroRoutes(router *gin.Engine, db *gorm.DB) *gin.RouterGroup {
	repo := repository.NewGeneroRepository(db)
	resource := resource.NewGeneroResource(repo)

	group := router.Group("/generos")
	{
		group.POST("/", resource.Create)
		group.GET("/", resource.GetAll)
		group.PUT("/:id", resource.Update)
		group.DELETE("/:id", resource.Delete)
	}
	return group
}
