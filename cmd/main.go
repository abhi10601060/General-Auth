package main

import (
	"fmt"
	"synapse/auth/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi From Auth Service")

	r := gin.Default()

	r.GET("/ping", handler.Pong)
	r.POST("/signup", handler.SignUp)
	r.POST("/login", handler.Login)

	r.Run(":8000")
}
