package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/mailing/models"
	service "github.com/wyllisMonteiro/mailing/service"
)

// GetCampaign : Return JSON of a campaign or error
func GetCampaign(w http.ResponseWriter, req *http.Request) {

	fmt.Println("c le début le sang")

	urlParams := mux.Vars(req)

	getCampaign, err := models.CampaignFindByID(urlParams["id"])
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de récupérer la campagne")
		return
	}

	fmt.Println("J'envoie l'id le sang")
	service.SendIDCampaign(getCampaign.ID)
	service.WriteJSON(w, http.StatusOK, getCampaign)

}

// CreateCampaign : Return JSON of created campaign or error
func CreateCampaign(w http.ResponseWriter, req *http.Request) {
	var body models.CreateCampaignRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de créer la campagne")
		return
	}

	createCampaign, err := models.CreateCampaign(w, body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de créer la campagne")
		return
	}

	service.SendIDCampaign(createCampaign.ID)

	service.WriteJSON(w, http.StatusOK, createCampaign)
}
