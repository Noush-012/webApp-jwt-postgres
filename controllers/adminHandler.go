package controllers

import (
	"fmt"
	"net/http"

	"github.com/Noush-012/web_jwt/helpers"
	"github.com/gin-gonic/gin"
)

// ================================== ADMIN LOGIN SECTION ================================== //

// To render admin login page
func AdminLogin(c *gin.Context) {
	fmt.Println("login admin")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	c.HTML(http.StatusOK, "adminLogin.html", "")
}

// Post admin login
func AdminLoginSubmit(c *gin.Context) {
	fmt.Println("Admin trying to login")

	// validate user if success return ID else error
	userVal, ok := helpers.AdminValidation(struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}{
		Email:    c.Request.PostFormValue("adminid"),
		Password: c.Request.PostFormValue("password"),
	})

	if !ok {
		c.Redirect(http.StatusSeeOther, "/admin")
		return
	}
	// If admin valid generate JWT and set into cookie. cookie name is set to"admin"
	if !helpers.JwtCookieSetup(c, "admin", userVal) {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/home")

}

// ================================== ADMIN HOME SECTION ================================== //

// To render home page
func AdminHome(c *gin.Context) {
	fmt.Println("Admin Home Page")

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.HTML(http.StatusOK, "adminHome.html", nil)
}

// To logout admin
func LogoutAdmin(c *gin.Context) {
	fmt.Println("User Logged out")

	_, ok := helpers.GetCookieVal(c, "admin")

	if !ok {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.SetCookie("admin", "", -1, "", "", false, true)

	//atlast redirect to login page
	c.Redirect(http.StatusSeeOther, "/admin")
}
