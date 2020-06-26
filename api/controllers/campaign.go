package controllers

import (
  "net/http"
  "fmt"
  "encoding/json"
	campaign "github.com/wyllisMonteiro/mailing/api/repositories/campaign"
	service "github.com/wyllisMonteiro/mailing/api/service"
)

/**
  * Post broadcast list
  * 
  * 	POST /campaign
  * 	req.body :
  *		{
  *	  		message   			  string
  *		}
  * 
  */
func Campaign(w http.ResponseWriter, req *http.Request) {

  if req.Method == "POST" {
    CreateCampaign(w, req)
  }
  
  /*
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

  service.WriteJSON(w, http.StatusOK, broad)*/
}

func CreateCampaign(w http.ResponseWriter, req *http.Request) {
  var body campaign.CreateCampaignRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
    fmt.Println("Error")
    return
	}

  createCampaign, err := campaign.CreateCampaign(w, body)
  if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de créer la campagne")
    return
  }

  service.WriteJSON(w, http.StatusOK, createCampaign)

}