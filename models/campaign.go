package models

import (
	"fmt"
	"net/http"
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

	res, err := DB.Exec("INSERT `campaign`(`message`, `broadcast_id`) VALUES (?, ?)", createCampaignRequest.Message, broad.ID)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(3)
		return createCampaignResponse, err
	}

	campaignID, err := res.LastInsertId()
	if err != nil {
		return createCampaignResponse, err
	}

	createCampaignResponse.ID = campaignID
	createCampaignResponse.Message = createCampaignRequest.Message
	createCampaignResponse.BroadcastID = string(broad.ID)

	return createCampaignResponse, nil
}

// CampaignFindByID : Get campaign according to ID and return campaign
func CampaignFindByID(campaignID string) (CreateCampaignResponse, error) {
	var createCampaignResponse CreateCampaignResponse = CreateCampaignResponse{}

	err := DB.QueryRow("SELECT * FROM campaign WHERE id = ?",
		campaignID).Scan(&createCampaignResponse.ID,
		&createCampaignResponse.Message,
		&createCampaignResponse.BroadcastID)

	if err != nil {
		return createCampaignResponse, err
	}

	return createCampaignResponse, nil
}
