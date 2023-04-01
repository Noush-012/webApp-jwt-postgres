package routes

import (
	"github.com/Noush-012/web_jwt/controllers"
	"github.com/Noush-012/web_jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	// User signup & submit routes
	r.GET("/signup", controllers.UserSignup)
	r.POST("/signup", controllers.SignupSubmition)

	// User login & submit routes
	r.GET("/", controllers.LoginPage)
	r.POST("/", controllers.UserLoginSubmission)

	// User home and logout routes
	r.GET("/home", middleware.UserAuthentiaction, controllers.UserHome)
	r.POST("/logout", controllers.LogoutUser)

}
