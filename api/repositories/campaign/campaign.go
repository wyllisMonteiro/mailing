package campaign

import (
	"net/http"
	"github.com/wyllisMonteiro/mailing/api/service"
	broadcast "github.com/wyllisMonteiro/mailing/api/repositories/broadcast"
	config "github.com/wyllisMonteiro/mailing/api/config"
)

type CreateCampaignRequest struct {
  Message string `json:"message"`
  BroadcastName string `json:"broadcastName"`
}

type CreateCampaignResponse struct {
	ID int64
	Message string
	BroadcastID int
}

func CreateCampaign(w http.ResponseWriter, createCampaignRequest CreateCampaignRequest) (CreateCampaignResponse, error) {

	var createCampaignResponse CreateCampaignResponse = CreateCampaignResponse{}

	broad, err := broadcast.FindBy("name", createCampaignRequest.BroadcastName)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la liste de diffusion n'a pas été effectué")
		return createCampaignResponse, err
	}

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	res, err := db.Exec("INSERT `campaign`(`message`, `broadcast_id`) VALUES (?, ?)", createCampaignRequest.Message, broad.ID)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la campagne n'a pas été effectué")
		return createCampaignResponse, err
	}

	campaign_id, err := res.LastInsertId()
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, mail introuvable")
		return createCampaignResponse, err
	}

	createCampaignResponse.ID = campaign_id
	createCampaignResponse.Message = createCampaignRequest.Message
	createCampaignResponse.BroadcastID = broad.ID

	return createCampaignResponse, nil
}