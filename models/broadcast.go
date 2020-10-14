package models

import (
	"log"
	"strconv"
)

// Broadcasts : Structure format
type Broadcasts struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Mails       []string `json:"mails"`
}

// Broadcast : Structure format in db
type Broadcast struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Mail        string `json:"mail"`
}

// BroadcastFindBy : return broadcast according to params function
func BroadcastFindBy(key string, val string) (Broadcast, error) {
	var broadResponse Broadcast

	err := DB.QueryRow("SELECT id, name, description FROM broadcast WHERE "+key+" = ?", val).Scan(&broadResponse.ID, &broadResponse.Name, &broadResponse.Description)

	if err != nil {
		log.Println(err)
		return broadResponse, err
	}

	return broadResponse, nil
}

// BroadcastFindWithSubs : return broadcast with subscriber mail
func BroadcastFindWithSubs(name string) (Broadcasts, error) {
	var broadsResponse Broadcasts
	var broadResponse Broadcast

	query := `
		SELECT
			broadcast.id,
			broadcast.name,
			broadcast.description,
			subscriber.mail
		FROM 
			broadcast,
			broadcast_subscriber,
			subscriber
		WHERE broadcast.name = ?
		AND broadcast.id = broadcast_subscriber.broadcast_id
		AND subscriber.id = broadcast_subscriber.subscriber_id
	`

	results, err := DB.Query(query, name)
	if err != nil {
		return broadsResponse, err
	}

	for results.Next() {
		err = results.Scan(&broadResponse.ID, &broadResponse.Name, &broadResponse.Description, &broadResponse.Mail)
		if err != nil {
			return broadsResponse, err
		}

		broadsResponse.Mails = append(broadsResponse.Mails, broadResponse.Mail)
	}

	broadsResponse.ID = broadResponse.ID
	broadsResponse.Name = broadResponse.Name
	broadsResponse.Description = broadResponse.Description

	return broadsResponse, nil
}

// CreateBroadcast : Create broadcast and return broadcast created
func CreateBroadcast(createBroadcastRequest Broadcasts) (Broadcasts, error) {
	res, err := DB.Exec("INSERT `broadcast`(`name`, `description`) VALUES (?, ?)", createBroadcastRequest.Name, createBroadcastRequest.Description)
	if err != nil {
		return createBroadcastRequest, err
	}

	broadcastID, err := res.LastInsertId()
	if err != nil {
		return createBroadcastRequest, err
	}

	createBroadcastRequest.ID = strconv.Itoa(int(broadcastID))

	for mailIndex := 0; mailIndex < len(createBroadcastRequest.Mails); mailIndex++ {
		subscriber, err := SubscriberFindBy("mail", createBroadcastRequest.Mails[mailIndex])
		if err != nil {
			return createBroadcastRequest, err
		}

		insert, err := DB.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broadcastID, subscriber.ID)

		if err != nil {
			return createBroadcastRequest, err
		}

		defer insert.Close()
	}

	return createBroadcastRequest, nil
}

// SubRequest : Structure format of request
type SubRequest struct {
	BroadcastName  string `json:"broadcast_name"`
	SubscriberMail string `json:"subscriber_mail"`
}

// BroadcastAddSubscriber : Add subscriber to broadcast and return request called
func BroadcastAddSubscriber(subRequest SubRequest) (SubRequest, error) {
	subscriber, err := SubscriberFindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		return subRequest, err
	}

	broad, err := BroadcastFindBy("name", subRequest.BroadcastName)
	if err != nil {
		return subRequest, err
	}

	insert, err := DB.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broad.ID, subscriber.ID)

	if err != nil {
		return subRequest, err
	}

	defer insert.Close()

	return subRequest, nil
}

// BroadcastDeleteSubscriber : delete subscriber to broadcast and return request called
func BroadcastDeleteSubscriber(subRequest SubRequest) (SubRequest, error) {
	subscriber, err := SubscriberFindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		return subRequest, err
	}

	broad, err := BroadcastFindBy("name", subRequest.BroadcastName)
	if err != nil {
		return subRequest, err
	}

	delete, err := DB.Query("DELETE FROM `broadcast_subscriber` WHERE broadcast_id = ? AND subscriber_id = ?", broad.ID, subscriber.ID)

	if err != nil {
		return subRequest, err
	}

	defer delete.Close()

	return subRequest, nil
}
