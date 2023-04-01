package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

// To load environment variables
func LoadENV() {

	if godotenv.Load() != nil {
		fmt.Println("Faild to load env")
		return
	}
	fmt.Println("Successfully loaded env")

}
