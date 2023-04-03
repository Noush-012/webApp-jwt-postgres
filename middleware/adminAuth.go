package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Noush-012/web_jwt/controllers"
	"github.com/Noush-012/web_jwt/helpers"
	"github.com/Noush-012/web_jwt/initializers"
	"github.com/Noush-012/web_jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Admin authentication

func AdminAuthentication(c *gin.Context) {
	fmt.Println("Processing user authentication...")

	token, ok := helpers.GetToken(c, "admin")
	if !ok {
		c.Abort()
		controllers.AdminLogin(c)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //valid token

		// Check token time expired or not
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			fmt.Println("the token is timeout")
			c.Abort()
			c.Redirect(http.StatusSeeOther, "/admin")
			return
		}
		// TO cross check whether admin credentials match in database
		var admin models.Admin

		adminID := uint(claims["userId"].(float64))

		initializers.DB.Find(&admin, "id = ?", adminID)
		// fmt.Println("Got email id from db:", admin.Email)
		if admin.ID == 0 {
			c.Abort()
			controllers.AdminLogin(c)
			// controllers.AdminHome(c)
		}
		// Setting admin ID in ctx
		c.Set("adminID", adminID)

		// Render admin home page if authentication successful
		if c.Request.URL.Path == "/admin" {
			c.Abort()
			c.Redirect(http.StatusSeeOther, "/admin/home")
			return
		}
		c.Next()

		fmt.Println("Admin athentication finished")
	}

}
