package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MySigningKey : Allow to generate JWT with a secret key
var MySigningKey = []byte("mysupersecret")
var tokenString string

// GenerateJWT : Generate a token using JWT
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Wyllis Monteiro"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(MySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetToken : Getter for token
func GetToken() string {
	return tokenString
}
