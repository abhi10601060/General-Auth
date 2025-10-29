package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading env file : ", err.Error())
	}
	fmt.Println("env file loaded successfully")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}