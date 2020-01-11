package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spaceCh1mp/pow/server/pkg/group"
	"github.com/spaceCh1mp/pow/server/pkg/transaction"
	"github.com/spaceCh1mp/pow/server/pkg/user"
	"github.com/spaceCh1mp/pow/server/routes"
)

func main() {

	r := gin.Default()

	//starts up the server.
	p := &routes.Res{
		BuildPath: "./client/build",
		NotFound:  "./assets/404/index.html",
	}

	//define services to run
	services := []func() error{
		user.Config,
		group.Config,
		transaction.Config,
	}

	//run services on seperate threads
	for _, v := range services {
		go v()
	}

	//register router
	routes.Router(r, p)

	r.Run(":5005")
}
