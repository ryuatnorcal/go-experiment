package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Example struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
}

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	server.GET("/params/:id/:slug", func(c *gin.Context) {
		id := c.Param("id")
		slug := c.Param("slug")
		fmt.Println("id", id)
		c.JSON(200, gin.H{
			"data": gin.H{
				"id":   id,
				"slug": slug,
			},
		})
	})

	server.POST("/post", func(c *gin.Context) {
		var json Example
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		id := json.Id
		slug := json.Slug
		fmt.Println("id", id)
		c.JSON(200, gin.H{
			"data": gin.H{
				"id":   id,
				"slug": slug,
			},
		})
	})
}
