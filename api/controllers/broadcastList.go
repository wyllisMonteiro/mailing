package controllers

import (
	//"github.com/wyllisMonteiro/mailing/api/service"
	"net/http"
	"fmt"
	"encoding/json"
)

type PostBody struct {
	Name   string    `json:"name"`
    Description string `json:"description"`
    Mails []string `json:"mails"`
}

func BroadCastList(w http.ResponseWriter, req *http.Request) {

	var body PostBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("error")
	}

}