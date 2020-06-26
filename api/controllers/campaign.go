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

  if req.Method == "GET" {
    GetCampaign(w, req)
  }

  if req.Method == "POST" {
    CreateCampaign(w, req)
  }

}

func GetCampaign(w http.ResponseWriter, req *http.Request) {
  var body campaign.GetCampaignRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
    fmt.Println("Error")
    return
  }
  
  getCampaign, err := campaign.FindByID(body.ID)
  if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de créer la campagne")
    return
  }

  //service.SendIdCampaign(createCampaign.ID)
  service.WriteJSON(w, http.StatusOK, getCampaign)

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

  //service.SendIdCampaign(createCampaign.ID)
  service.WriteJSON(w, http.StatusOK, createCampaign)
}