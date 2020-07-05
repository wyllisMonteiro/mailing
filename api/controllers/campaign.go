package controllers

import (
  "net/http"
  "fmt"
  "encoding/json"
	repo "github.com/wyllisMonteiro/mailing/api/repositories"
  service "github.com/wyllisMonteiro/mailing/api/service"
	"github.com/gorilla/mux"
)

func GetCampaign(w http.ResponseWriter, req *http.Request) {

  urlParams := mux.Vars(req)
  
  getCampaign, err := repo.CampaignFindByID(urlParams["id"])
  if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de créer la campagne")
    return
  }

  //service.SendIdCampaign(createCampaign.ID)
  service.WriteJSON(w, http.StatusOK, getCampaign)

}

func CreateCampaign(w http.ResponseWriter, req *http.Request) {
  var body repo.CreateCampaignRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
    fmt.Println("Error")
    return
	}

  createCampaign, err := repo.CreateCampaign(w, body)
  if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de créer la campagne")
    return
  }

  //service.SendIdCampaign(createCampaign.ID)
  service.WriteJSON(w, http.StatusOK, createCampaign)
}