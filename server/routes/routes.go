package routes

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

//Res sets the path to resources needed to run the server
type Res struct {
	BuildPath string
	NotFound  string
}

//Router defines the routes for the application
func Router(r *gin.Engine, p *Res) {
	ro := []struct {
		name string
	}{
		{"/"},
		{"/login"},
		{"/register"},
		{"/profile"},
		{"/dashboard"},
	}

	//TODO
	//Evaluate if user is logged in
	//redirect to dashboard with user details if true
	//redirect to login if false

	for _, er := range ro {
		r.Use(static.Serve(er.name, static.LocalFile(p.BuildPath, true)))
	}

	//Handle error 404
	r.LoadHTMLFiles(p.NotFound)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "index.html", gin.H{})
	})
}
