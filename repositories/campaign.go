package repositories

import (
	"net/http"

	config "github.com/wyllisMonteiro/mailing/api/config"
)

type CreateCampaignRequest struct {
	Message       string `json:"message"`
	BroadcastName string `json:"broadcastName"`
}

type CreateCampaignResponse struct {
	ID          int64  `json:"id"`
	Message     string `json:"message"`
	BroadcastID string `json:"broadcast_id"`
}

func CreateCampaign(w http.ResponseWriter, createCampaignRequest CreateCampaignRequest) (CreateCampaignResponse, error) {

	var createCampaignResponse CreateCampaignResponse = CreateCampaignResponse{}

	broad, err := BroadcastFindBy("name", createCampaignRequest.BroadcastName)
	if err != nil {
		return createCampaignResponse, err
	}

	db, err := config.ConnectToBDD()
	if err != nil {
		return createCampaignResponse, err
	}

	defer db.Close()

	res, err := db.Exec("INSERT `campaign`(`message`, `broadcast_id`) VALUES (?, ?)", createCampaignRequest.Message, broad.ID)
	if err != nil {
		return createCampaignResponse, err
	}

	campaign_id, err := res.LastInsertId()
	if err != nil {
		return createCampaignResponse, err
	}

	createCampaignResponse.ID = campaign_id
	createCampaignResponse.Message = createCampaignRequest.Message
	createCampaignResponse.BroadcastID = string(broad.ID)

	return createCampaignResponse, nil
}

func CampaignFindByID(campaignId string) (CreateCampaignResponse, error) {
	var createCampaignResponse CreateCampaignResponse = CreateCampaignResponse{}

	db, err := config.ConnectToBDD()
	if err != nil {
		return createCampaignResponse, err
	}

	defer db.Close()

	err = db.QueryRow("SELECT * FROM campaign WHERE id = ?",
		campaignId).Scan(&createCampaignResponse.ID,
		&createCampaignResponse.Message,
		&createCampaignResponse.BroadcastID)

	if err != nil {
		return createCampaignResponse, err
	}

	return createCampaignResponse, nil
}
