package main

import (
	"github.com/harcokun/go-crud/initializers"
	"github.com/harcokun/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToTB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
