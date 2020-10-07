package repositories

import (
	"fmt"
	"net/http"

	config "github.com/wyllisMonteiro/mailing/config"
)

// CreateCampaignRequest : Structure format
type CreateCampaignRequest struct {
	Message       string `json:"message"`
	BroadcastName string `json:"broadcastName"`
}

// CreateCampaignResponse : Structure format in db
type CreateCampaignResponse struct {
	ID          int64  `json:"id"`
	Message     string `json:"message"`
	BroadcastID string `json:"broadcast_id"`
}

// CreateCampaign : Create campaign and return campaign
func CreateCampaign(w http.ResponseWriter, createCampaignRequest CreateCampaignRequest) (CreateCampaignResponse, error) {
	var createCampaignResponse CreateCampaignResponse = CreateCampaignResponse{}

	broad, err := BroadcastFindBy("name", createCampaignRequest.BroadcastName)
	if err != nil {
		fmt.Println(err.Error())
		return createCampaignResponse, err
	}

	db, err := config.ConnectToBDD()
	if err != nil {
		fmt.Println(err.Error())
		return createCampaignResponse, err
	}

	defer db.Close()

	res, err := db.Exec("INSERT `campaign`(`message`, `broadcast_id`) VALUES (?, ?)", createCampaignRequest.Message, broad.ID)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(3)
		return createCampaignResponse, err
	}

	campaign_id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(4)
		return createCampaignResponse, err
	}

	createCampaignResponse.ID = campaign_id
	createCampaignResponse.Message = createCampaignRequest.Message
	createCampaignResponse.BroadcastID = string(broad.ID)

	return createCampaignResponse, nil
}

// CampaignFindByID : Get campaign according to ID and return campaign
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
