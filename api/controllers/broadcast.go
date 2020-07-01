package controllers

import (
	"fmt"
	"net/http"
  "encoding/json"
  service "github.com/wyllisMonteiro/mailing/api/service"
	repos "github.com/wyllisMonteiro/mailing/api/repositories"
)

/**
  * Broadcast controller
  * 
  * 	GET | POST
  * 
  */
func BroadCast(w http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    GetBroadcast(w, req)
  }
  
  if req.Method == "POST" {
    CreateBroadcast(w, req)
  }
}

/**
  * GET /broadcast
  * 	req.body :
  *		{
  *	  	name   			  string
  *		}
  */
  func GetBroadcast(w http.ResponseWriter, req *http.Request) {
    var body repos.GetBroadcastRequest
  
    err := json.NewDecoder(req.Body).Decode(&body)
    if err != nil {
      fmt.Println("Error")
    }
  
    broad, err := repos.BroadcastFindWithSubs(body.Name)
    if err != nil {
      fmt.Println(err.Error())
      service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, liste de diffusion introuvable")
      return
    }

    service.WriteJSON(w, http.StatusOK, broad)
  }

/**
  * POST /broadcast
  * 	req.body :
  *		{
  *	  	name   			  string
  * 		description 	string
  *  		mails 			[]string
  *		}
  */
func CreateBroadcast(w http.ResponseWriter, req *http.Request) {
  var body repos.CreateBroadcastRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	repos.CreateBroadcast(w, body)
}

/**
  * Add subscriber in broadcast list
  * 
  * 	POST /broadcast/add/subscriber
  * 	req.body :
  *		{
  *	  	broadcastName 		string
  *			subscriberMail 		string
  *		}
  * 
  */
func AddSubscriber(w http.ResponseWriter, req *http.Request) {
	var body repos.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	repos.BroadcastAddSubscriber(w, body)
}

/**
  * Delete subscriber in broadcast list
  * 
  * 	DELETE /broadcast/delete/subscriber
  * 	req.body :
  *		{
  *	  	broadcastName 		string
  *			subscriberMail 		string
  *		}
  * 
  */
func DeleteSubscriber(w http.ResponseWriter, req *http.Request) {
	var body repos.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error")
	}

	repos.BroadcastDeleteSubscriber(w, body)
}