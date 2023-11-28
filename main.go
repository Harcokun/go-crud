package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/harcokun/go-crud/initializers"
	"github.com/harcokun/go-crud/models"
	routers "github.com/harcokun/go-crud/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToTB()
	initializers.DB.Debug().AutoMigrate(&models.Post{})
}

func main() {
	// router := gin.Default()

	router := gin.New()
	// gin.SetMode(gin.ReleaseMode) // uncomment this line to run in production mode
	api := router.Group("/api/v1")
	routers.SetRoutes(api)

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	go router.Run(port)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
