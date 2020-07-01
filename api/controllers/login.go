package controllers

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/api/service"
	repo "github.com/wyllisMonteiro/mailing/api/repositories"
	"net/http"
	"log"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user, err := repo.ClientFindBy("mail", "wyllismonteiro@gmail.com")
	if err != nil {
		panic(err.Error())
		return
	}

	match, err := service.ComparePasswordAndHash("w", user.Password)
  	if err != nil {
		log.Fatal(err)
		return
	}

	if !match {
		return
	}

	validToken, err := service.GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	repo.UpdateToken(validToken, user.ID)
	_, _ = fmt.Fprintf(w, validToken)
}