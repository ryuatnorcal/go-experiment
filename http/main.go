package main

import (
	"example.com/m/v2/api"
	"example.com/m/v2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	api.RegisterAPIRoutes(server)
	server.Run(":8080")
}
