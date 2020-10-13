package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wyllisMonteiro/mailing/models"
	service "github.com/wyllisMonteiro/mailing/service"
)

// GetBroadcast : Return JSON of all broadcast or error
func GetBroadcast(w http.ResponseWriter, req *http.Request) {

	broadcastName := req.FormValue("name")

	broad, err := models.BroadcastFindWithSubs(broadcastName)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, liste de diffusion introuvable")
		return
	}

	service.WriteJSON(w, http.StatusOK, broad)
}

// CreateBroadcast : Return JSON of created broadcast or error
func CreateBroadcast(w http.ResponseWriter, req *http.Request) {
	var body models.Broadcasts

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue")
		return
	}

	broad, err := models.CreateBroadcast(body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la liste de diffusion a échoué")
		return
	}

	service.WriteJSON(w, http.StatusOK, broad)
}

// AddSubscriber : Return JSON of added subscriber to broadcast or error
func AddSubscriber(w http.ResponseWriter, req *http.Request) {
	var body models.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, l'ajout d'un subscriber a échoué")
		return
	}

	subRequest, err := models.BroadcastAddSubscriber(body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, l'ajout d'un subscriber a échoué")
		return
	}

	service.WriteJSON(w, http.StatusOK, subRequest)

}

// DeleteSubscriber : Return JSON of deleted subscriber to broadcast or error
func DeleteSubscriber(w http.ResponseWriter, req *http.Request) {
	var body models.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la supression d'un subscriber a échoué")
		return
	}

	subRequest, err := models.BroadcastDeleteSubscriber(body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la supression d'un subscriber a échoué")
		return
	}

	service.WriteJSON(w, http.StatusOK, subRequest)
}
