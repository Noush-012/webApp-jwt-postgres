package controllers

import (
	"fmt"
	"net/http"

	"github.com/Noush-012/web_jwt/helpers"
	"github.com/Noush-012/web_jwt/initializers"
	"github.com/Noush-012/web_jwt/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// ================================== ADMIN LOGIN SECTION ================================== //

var AdminTempMessage interface{}

// To render admin login page
func AdminLogin(c *gin.Context) {
	fmt.Println("login admin")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	data := gin.H{
		"Color": "text-danger",
		"Alert": AdminTempMessage,
	}
	c.HTML(http.StatusOK, "adminLogin.html", data)
	AdminTempMessage = nil
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
		AdminTempMessage = userVal
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

	// Record array to hold all user value from DB
	var record []models.User
	initializers.DB.Find(&record)
	// To store user details and link to template
	type field struct {
		ID        int
		UserId    uint
		FirstName string
		LastName  string
		Email     string
		Status    bool
	}

	//slice to store all user
	var arrayOfField []field
	for i, v := range record {
		arrayOfField = append(arrayOfField, field{
			ID:        i + 1,
			UserId:    v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Status:    v.Status,
		})
	}
	c.HTML(http.StatusOK, "adminHome.html", arrayOfField)
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

// ================================== ADMIN PRIVILEGE SECTION ================================== //

// Block user
func BlockUser(c *gin.Context) {
	fmt.Println("Admin tries to block")

	userId := c.Params.ByName("id")

	if c.Params.ByName("status") == "block" {
		initializers.DB.Model(&models.User{}).Where("id = ?", userId).Update("status", false)
	} else {
		initializers.DB.Model(&models.User{}).Where("id = ?", userId).Update("status", true)
	}
	c.Redirect(http.StatusSeeOther, "/admin/home")
}

// To delete
func DeleteUser(c *gin.Context) {

	userId := c.Param("id")

	initializers.DB.Clauses(clause.OnConflict{DoNothing: true}).Delete(&models.User{}, "id = ?", userId)

	c.Redirect(http.StatusSeeOther, "/admin/home")
}
