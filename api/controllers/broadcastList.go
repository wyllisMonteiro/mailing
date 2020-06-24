package controllers

import (
	//"github.com/wyllisMonteiro/mailing/api/service"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/wyllisMonteiro/mailing/api/repositories"
)

type PostBody struct {
	Name   string
    Description string
    Mails []string
}

func BroadCastList(w http.ResponseWriter, req *http.Request) {

	var body PostBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("yes")
	}

	repositories.AddBroadcastList(body)

	

}