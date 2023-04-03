package routes

import (
	"github.com/Noush-012/web_jwt/controllers"
	"github.com/Noush-012/web_jwt/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {

	// Admin login & submit routes
	r.GET("/admin", middleware.AdminAuthentication, controllers.AdminLogin)
	r.POST("/admin", controllers.AdminLoginSubmit)

	// Admin home & logout routes
	r.GET("/admin/home", middleware.AdminAuthentication, controllers.AdminHome)
	r.GET("/adminlogout", controllers.LogoutAdmin)
}
