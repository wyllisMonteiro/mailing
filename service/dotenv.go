package service

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// EnvFilePath : Location environment file
var EnvFilePath = "./.env"

// GoDotEnvVariable : Get variable value from key after loading .env file
func GoDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

// GetVarsMail : Get all variables to send mail
func GetVarsMail() (string, int, string, string, string, string, error) {
	host, err := GoDotEnvVariable("HOSTMAIL")
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	portStr, err := GoDotEnvVariable("PORTMAIL")
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	user, err := GoDotEnvVariable("USERMAIL")
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	pass, err := GoDotEnvVariable("PASSMAIL")
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	from, err := GoDotEnvVariable("FROMMAIL")
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	to, err := GoDotEnvVariable("TOMAIL")
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, "", "", "", "", err
	}

	return host, port, user, pass, from, to, nil
}
