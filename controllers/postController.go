package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/harcokun/go-crud/initializers"
	"github.com/harcokun/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	// Return a result

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return the result
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off the URL
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Return the result
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off the URL
	id := c.Param("id")

	// Get the id off the req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return the result
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off the URL
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Delete it
	result := initializers.DB.Delete(&models.Post{}, id)

	// Return the result
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "Post is deleted",
	})
}
