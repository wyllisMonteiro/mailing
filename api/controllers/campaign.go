package controllers

import (
  "net/http"
  "fmt"
  "encoding/json"
	config "github.com/wyllisMonteiro/mailing/api/config"
	broadcast "github.com/wyllisMonteiro/mailing/api/repositories/broadcast"
	service "github.com/wyllisMonteiro/mailing/api/service"
)

/**
  * Post broadcast list
  * 
  * 	POST /campaign
  * 	req.body :
  *		{
  *	  		name   			  string
  *		}
  * 
  */
func Campaign(w http.ResponseWriter, req *http.Request) {
  var body broadcast.GetBroadcastRequest
  
  err := json.NewDecoder(req.Body).Decode(&body)
  if err != nil {
    fmt.Println("Error")
  }

  broad, err := broadcast.FindWithSubs(body.Name)
  if err != nil {
    fmt.Println(err.Error())
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, liste de diffusion introuvable")
    return
  }

  db, err := config.ConnectToBDD()
  if err != nil {
    return
  }

  service.RabbitSend()

  insert, err := db.Query("INSERT `mail`(`description`) VALUES (?)", broad.Description)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la liste de diffusion n'a pas été effectué")
		return 
  }
  
  defer insert.Close()

  service.WriteJSON(w, http.StatusOK, broad)
}