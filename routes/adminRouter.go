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
	r.GET("/admin/logout", controllers.LogoutAdmin)

	// Admin privilege routes
	// Block
	r.GET("/admin/blockuser/:status/:id", controllers.BlockUser)
	// Delete
	r.GET("/admin/deleteuser/:id", controllers.DeleteUser)
	// Add user
	r.GET("/admin/adduser", middleware.AdminAuthentication, controllers.AdminAddUser)
	r.POST("/admin/adduser", middleware.AdminAuthentication, controllers.PostAddUserAdmin)
}
