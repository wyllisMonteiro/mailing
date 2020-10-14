package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/mailing/models"
	"github.com/wyllisMonteiro/mailing/service"
)

// GetCampaign : Return JSON of a campaign or error
func GetCampaign(w http.ResponseWriter, req *http.Request) {

	fmt.Println("On est dans le controller")

	urlParams := mux.Vars(req)
	id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de récupérer la campagne")
		return
	}

	getCampaign, err := models.CampaignFindByID(id)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Impossible de récupérer la campagne")
		return
	}

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
