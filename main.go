package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	api := router.Group("/api")
	api.GET("/", func(g *gin.Context) {
		g.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":5000")
}
