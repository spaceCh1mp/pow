package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//serve react build
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	router.Run(":5000")
}
