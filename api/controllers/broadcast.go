package controllers

import (
	"fmt"
	"net/http"
  "encoding/json"
  service "github.com/wyllisMonteiro/mailing/api/service"
  repos "github.com/wyllisMonteiro/mailing/api/repositories"
)

func GetBroadcast(w http.ResponseWriter, req *http.Request) {

  broadcastName := req.FormValue("name")
  
  broad, err := repos.BroadcastFindWithSubs(broadcastName)
  if err != nil {
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, liste de diffusion introuvable")
    return
  }

  service.WriteJSON(w, http.StatusOK, broad)
}

func CreateBroadcast(w http.ResponseWriter, req *http.Request) {
  var body repos.Broadcasts

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue")
    return
  }

  broad, err := repos.CreateBroadcast(body)
  if err != nil {
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la liste de diffusion a échoué")
    return
  }

  service.WriteJSON(w, http.StatusOK, broad)
}

func AddSubscriber(w http.ResponseWriter, req *http.Request) {
	var body repos.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, l'ajout d'un subscriber a échoué")
    return
	}

  subRequest, err := repos.BroadcastAddSubscriber(body)
  if err != nil {
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, l'ajout d'un subscriber a échoué")
    return
  }

	service.WriteJSON(w, http.StatusOK, subRequest)
  
}

func DeleteSubscriber(w http.ResponseWriter, req *http.Request) {
	var body repos.SubRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la supression d'un subscriber a échoué")
    return
  }

  subRequest, err := repos.BroadcastDeleteSubscriber(body)
	if err != nil {
    service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la supression d'un subscriber a échoué")
    return
  }

	service.WriteJSON(w, http.StatusOK, subRequest)
}