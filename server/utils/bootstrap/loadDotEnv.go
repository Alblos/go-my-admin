package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
		return err
	}
	return nil
}
