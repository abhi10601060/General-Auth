package token

import (
	"log"
	"synapse/auth/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret_key = []byte("Synapse_Rocks")
)

type claims struct {
	Id string
	*jwt.RegisteredClaims
}

func CreateJwtToken(user *model.User) (string , error){
	claims := &claims{
		Id: user.Id,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(1000 * time.Hour)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(secret_key)
	if err != nil {
		log.Println("error during createJwtToken : ", err.Error())
		return  "Not able to create token", err
	}

	return tokenStr, nil
}