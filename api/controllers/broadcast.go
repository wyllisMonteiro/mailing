package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	broadcast "github.com/wyllisMonteiro/mailing/api/repositories/broadcast"
)

/**
  * Get broadcast list
  * 
  * 	GET /broadcast
  * 	req.body :
  *		{
  *	  		name   			string
  * 		description 	string
  *  		mails 			[]string
  *		}
  * 
  */
func BroadCast(w http.ResponseWriter, req *http.Request) {
	var body broadcast.CreateBroadcastRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	broadcast.CreateBroadcast(w, body)
}

/**
  * Add subscriber in broadcast list
  * 
  * 	POST /broadcast/add/subscriber
  * 	req.body :
  *		{
  *	  		broadcastName 		string
  *			subscriberMail 		string
  *		}
  * 
  */
func AddSubscriber(w http.ResponseWriter, req *http.Request) {
	var body broadcast.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	broadcast.AddSubscriber(w, body)
}

/**
  * Delete subscriber in broadcast list
  * 
  * 	DELETE /broadcast/delete/subscriber
  * 	req.body :
  *		{
  *	  		broadcastName 		string
  *			subscriberMail 		string
  *		}
  * 
  */
func DeleteSubscriber(w http.ResponseWriter, req *http.Request) {
	var body broadcast.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	broadcast.DeleteSubscriber(w, body)
}