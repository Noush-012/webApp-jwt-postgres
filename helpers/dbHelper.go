package helpers

import (
	"fmt"
	"time"

	"github.com/Noush-012/web_jwt/initializers"
	"github.com/Noush-012/web_jwt/models"
)

// to delete JWT logs if the token time is expired
func DeleteBlackListToken() {

	initializers.DB.Where("end_time < ?", float64(time.Now().Unix())).Delete(&models.JwtSessionList{})

	fmt.Println("delted black listed token from database")
}
