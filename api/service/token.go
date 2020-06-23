package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var MY_SIGNING_KEY 	= []byte("mysupersecret")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Wyllis Monteiro"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(MY_SIGNING_KEY)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
