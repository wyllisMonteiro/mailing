package controllers

import (
	"encoding/json"
	"net/http"

	repo "github.com/wyllisMonteiro/mailing/api/repositories"
	"github.com/wyllisMonteiro/mailing/api/service"
)

func Login(w http.ResponseWriter, req *http.Request) {
	service.SendIdCampaign(20)

	var body repo.Client

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue")
		return
	}

	user, err := repo.ClientFindBy("mail", body.Mail)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, mauvais mail")
		return
	}

	match, err := service.ComparePasswordAndHash(body.Password, user.Password)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, mauvais mot de passe")
		return
	}

	if !match {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, mauvais mot de passe")
		return
	}

	validToken, err := service.GenerateJWT()
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, l'authentification a échoué")
		return
	}

	user.Token = validToken

	err = repo.UpdateToken(validToken, user.ID)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, le token n'a pas été ajouté à la base de données")
		return
	}

	service.WriteJSON(w, http.StatusOK, user)
}
