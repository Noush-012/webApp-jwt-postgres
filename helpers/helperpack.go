package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/Noush-012/web_jwt/initializers"
	"github.com/Noush-012/web_jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

// This function will perform signup validation using inbuilt validator package
func ValidateSignup(form struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
}) (interface{}, bool) { // return error if any

	// Create validator instance
	validate := validator.New()

	if err := validate.Struct(form); err != nil {
		// Create a map for store error message
		TempMessage := map[string]string{}
		for _, er := range err.(validator.ValidationErrors) {
			TempMessage[er.Namespace()] = "Enter " + er.Namespace() + " properly"
		}
		return TempMessage, false
	}

	// Check if the user already exist or not
	var user models.User

	initializers.DB.First(&user, "email = ?", form.Email)

	if user.ID != 0 { // User already exists
		fmt.Println("User already exist")
		return map[string]string{"Alert": "User already exist"}, false

	}
	// Hash user password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
	if err != nil {
		fmt.Println("Hashing failed")
		return map[string]string{"Password": "Error"}, false
	}
	// No errors need to hash the pass and store the data to database
	initializers.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
		Password:  string(hashedPass),
		Status:    true,
	})
	return map[string]string{"Color": "text-success",
		"Alert": "Sucessfully Account Created You Can Login",
	}, true // Everyting ok
}

// This function will perform login validation using inbuilt validator package
// Returns error if any else returns user ID if success
func ValidateUserLogin(form struct {
	Email    string `validate:"required,email"` // Frontend validation
	Password string `validate:"required"`
}) (interface{}, bool) {
	// Performs validation using validator package
	validate := validator.New()

	if err := validate.Struct(form); err != nil {
		var templateMessage = map[string]string{}

		for _, er := range err.(validator.ValidationErrors) {
			templateMessage[er.Field()] = "Enter " + er.Field() + " properly"
		}
		return templateMessage, false
	}

	// Check user exists in database
	var user models.User
	initializers.DB.Find(&user, "email = ?", form.Email)

	if user.ID == 0 { // if user not found
		return map[string]string{
			"Alert": "You are not a registered user you can signup",
			"Color": "text-danger",
		}, false
	}
	return user.ID, true

}

// JWT token & cookie setup for session handling
func JwtCookieSetup(c *gin.Context, name string, userId interface{}) bool {
	cookieTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId, // Store logged user info in token
		"exp":    cookieTime,
	})

	// Generate signed JWT token using evn var of secret key
	if tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY"))); err == nil {

		// Set cookie with signed string if no error
		c.SetCookie(name, tokenString, 10*60, "", "", false, true)

		fmt.Println("JWT sign & set Cookie successful")
		return true
	}
	fmt.Println("Failed JWT setup")
	return false

}

// To get cookie from client
func GetCookieVal(ctx *gin.Context, name string) (string, bool) {

	if cookieVal, err := ctx.Cookie(name); err == nil {
		return cookieVal, true
	}

	fmt.Println("Failed to get cookie")
	return "", false
}

func GetToken(ctx *gin.Context, name string) (*jwt.Token, bool) {
	//Delete expired token from JWT session List
	DeleteBlackListToken()

	// get cookie
	cookieval, ok := GetCookieVal(ctx, name)

	if !ok { // problem to get cookie so return false
		return nil, false
	}
	// check the user in JWT session list
	var JwtListCheck models.JwtSessionList

	initializers.DB.Find(&JwtListCheck, "token_string = ?", cookieval)
	if JwtListCheck.ID != 0 {
		return nil, false //this user is in session list
	}

	// Parse cookie to get JWT token
	token, err := jwt.Parse(cookieval, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println("failed to parse the cookie to token")
		return nil, false
	}
	return token, true

}
