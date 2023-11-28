package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/harcokun/go-crud/controllers"
)

func SetPostRoutes(router *gin.RouterGroup) {
	rg := router.Group("/posts")

	rg.POST("/", controllers.PostsCreate)
	rg.PUT("/:id", controllers.PostsUpdate)
	rg.GET("/", controllers.PostsIndex)
	rg.GET("/:id", controllers.PostsShow)
	rg.DELETE("/:id", controllers.PostsDelete)
}
