package main

import (
	"fmt"
	"net/http"

	"github.com/Noush-012/web_jwt/initializers"
	"github.com/Noush-012/web_jwt/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	// initalization of all prerequisites functions
	initializers.LoadENV()
	initializers.ConnToDB()
	initializers.MigrateToDB()
}

func main() {

	// Instance of gin frame work
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	// Calling routes for User & Admin using routes package
	// routes.Admin(r)
	routes.UserRoutes(r)

	//if invalid url found then show user login page
	r.NoRoute(func(ctx *gin.Context) {

		fmt.Println(ctx.Request.Method, "method")

		fmt.Println("Redirected login page")
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":3000")

}
