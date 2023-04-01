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

func UserAuthentiaction(ctx *gin.Context) {
	fmt.Println("user auth")

	token, ok := helpers.GetToken(ctx, "user")

	if !ok {
		fmt.Println("token not ok")
		if ctx.Request.URL.Path == "/signup" {
			ctx.Next()
			return
		}
		//else case abort and go to login page
		ctx.Abort()
		controllers.LoginPage(ctx)
		return
	}

	//if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check its time is not over
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println(" checking exp not ok")
			ctx.Abort()
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		//get the user from database using token claims
		userId := uint(claims["userId"].(float64))

		var user models.User
		initializers.DB.Find(&user, "id = ?", userId)

		if user.ID == 0 || !user.Status { //user not found or user blocked by admin
			fmt.Println("user not ok")
			//check the path that user want to signup login after he is not a valid user or blocked by admin

			if ctx.Request.URL.Path == "/signup" {
				fmt.Println("if user tries signup")
				ctx.Next()
				return
			}
			//any other path just show the login page
			ctx.Abort()
			fmt.Println("user not found but jwt is there admin deleted user")
			controllers.LoginPage(ctx)
			return
		}

		ctx.Set("userId", userId) //atach the user id in context if user is valid

		//if the user is valid and enter signup or login url show home page
		if ctx.Request.URL.Path == "/" || ctx.Request.URL.Path == "/signup" {
			fmt.Println("if user tries signup &login")
			ctx.Abort()
			ctx.Redirect(http.StatusSeeOther, "/home")
			return
		}

		//if all condition completed and the url is for home page
		ctx.Next()
	} else {
		fmt.Println("if token invalid ")
		//if the token is invalid or cant claim then show login page
		ctx.Abort()
		ctx.Redirect(http.StatusSeeOther, "/")
	}

}
