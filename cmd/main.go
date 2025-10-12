package main

import (
	"fmt"
	"synapse/auth/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)


func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading env file : ", err.Error())
	}
}

func main() {
	fmt.Println("Hi From Synapse Auth Service")

	r := gin.Default()

	r.GET("/ping", handler.Pong)
	r.POST("/signup" , handler.SignUp)
	r.POST("/login" , handler.Login)

	r.Run(":8000")
}
