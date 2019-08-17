package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spaceCh1mp/pow/server/routes"
)

func main() {

	r := gin.Default()
	p := &routes.Res{
		BuildPath: "./client/build",
		NotFound:  "./assets/404/index.html",
	}
	//register router
	routes.Router(r, p)

	r.Run(":5000")
}
