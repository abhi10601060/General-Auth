package handler

import (
	"log"
	"net/http"
	"synapse/auth/jwt"
	"synapse/auth/crypt"
	"synapse/auth/db"
	"synapse/auth/model"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Auth is Alive",
	})
}

func SignUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		log.Println("error in binding user is : ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	log.Println("received user is : ", user)

	if db.UserExist(&user) {
		c.JSON(409, gin.H{
			"message": "user exists already",
		})
		c.Abort()
		return
	}

	jwt_token, err := jwt.CreateJwtToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}

	salt, err := crypt.GenerateSalt(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}
	user.Salt = salt
	user.Password = crypt.HashPassword(user.Password, salt)

	res := db.AddUser(&user)
	if res == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Welcome to App",
		"token":   jwt_token,
	})
}

func Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		log.Println("error in binding user is : ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	log.Println("received user is : ", user)

	userFromDB, err := db.GetUserByID(user.Id)
	if err != nil {
		c.JSON(201, gin.H{
			"message": "Incorrect User Id",
		})
		c.Abort()
		return
	}

	if !crypt.VerifyPassword(user.Password, userFromDB.Salt, userFromDB.Password) {
		c.JSON(201, gin.H{
			"message": "Incorrect Credentials",
		})
		c.Abort()
		return
	}

	jwt_token, err := jwt.CreateJwtToken(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Welcome to App",
		"token":   jwt_token,
	})

}
