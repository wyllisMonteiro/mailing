package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
	broadcast "github.com/wyllisMonteiro/mailing/api/repositories/broadcast"
)

func BroadCast(w http.ResponseWriter, req *http.Request) {
	var body broadcast.CreateBroadcastRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	broadcast.CreateBroadcast(body)

}

func AddSubscriber(w http.ResponseWriter, req *http.Request) {
	var body broadcast.AddSubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	broadcast.AddSubscriber(body)
}