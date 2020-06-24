package controllers

import (
	//"github.com/wyllisMonteiro/mailing/api/service"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/wyllisMonteiro/mailing/api/repositories"
)

func BroadCastList(w http.ResponseWriter, req *http.Request) {

	var body repositories.PostBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	repositories.AddBroadcastList(body)

}