package controllers

import (
	"fmt"
	"net/http"

	"github.com/Noush-012/web_jwt/helpers"
	"github.com/Noush-012/web_jwt/initializers"
	"github.com/Noush-012/web_jwt/models"
	"github.com/gin-gonic/gin"
)

// To get any error message from any of the function as map key value pair
var FrontEndMessage interface{}

// ================================== SIGNUP SECTION ================================== //

// To render signup
func UserSignup(c *gin.Context) {
	fmt.Println("User on signup page...")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	c.HTML(http.StatusOK, "signup.html", FrontEndMessage)
	FrontEndMessage = nil // Changing value as nil for store another error for any function returns
}

// Post sgnup user
func SignupSubmition(c *gin.Context) {

	// Validate user data from signup form
	message, ok := helpers.ValidateSignup(struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		Email     string `validate:"required,email"`
		Password  string `validate:"required"`
	}{
		FirstName: c.Request.PostFormValue("fname"),
		LastName:  c.Request.PostFormValue("lname"),
		Email:     c.Request.PostFormValue("email"),
		Password:  c.Request.PostFormValue("password"),
	})
	// Show validation errors on signup page
	if !ok {
		fmt.Println("Form validation not ok!")
		FrontEndMessage = message
		UserSignup(c)
		return

	}
	// If validation success and also not an existing user then redirect to login page with success message
	FrontEndMessage = message
	c.Redirect(http.StatusSeeOther, "/")

}

// ================================== LOGIN SECTION ================================== //

// To render login
func LoginPage(c *gin.Context) {
	fmt.Println("login user")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	c.HTML(http.StatusOK, "userLogin.html", FrontEndMessage)
	FrontEndMessage = nil // Delete message after renders
}

// Post user login
func UserLoginSubmission(c *gin.Context) {
	fmt.Println("User trying to login")

	// validate user
	userVal, ok := helpers.ValidateUserLogin(struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}{
		Email:    c.Request.PostFormValue("email"),
		Password: c.Request.PostFormValue("password"),
	})

	// If any error shows error on front end
	if !ok {
		FrontEndMessage = userVal
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// If user valid generate JWT and set into cookie. cookie name is set to"user"
	if !helpers.JwtCookieSetup(c, "user", userVal) {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	c.Redirect(http.StatusSeeOther, "/home")

}

// ================================== USER HOME SECTION ================================== //

// To render home page
func UserHome(c *gin.Context) {
	fmt.Println("User on home page")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	var user models.User

	// Get user id from frontend and check in database
	if userId, ok := c.Get("userId"); ok {
		initializers.DB.Find(&user, "id = ?", userId)
		fmt.Println("User ID:", userId)
	}

	// Shows welcome user with name
	data := gin.H{
		"UserName": user.FirstName,
	}
	fmt.Println(user.FirstName)
	c.HTML(http.StatusOK, "userHome.html", data)
}

// To logout user
func LogoutUser(c *gin.Context) {
	fmt.Println("User Logged out")

	_, ok := helpers.GetCookieVal(c, "user")

	if !ok {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.SetCookie("user", "", -1, "", "", false, true)

	//atlast redirect to login page
	c.Redirect(http.StatusSeeOther, "/")
}
