package routers

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	SetPostRoutes(router)
}
