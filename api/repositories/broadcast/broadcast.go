package broadcast

import (
	"net/http"
	"github.com/wyllisMonteiro/mailing/api/config"
	sub "github.com/wyllisMonteiro/mailing/api/repositories/subscriber"
	"github.com/wyllisMonteiro/mailing/api/service"
)

type BroadcastResponse struct {
	ID   int    `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
}

/**
  * Get broadcast by id
  * 
  * 	in :
  *		key => field name in bdd you want to looking for
  *		val => field value in bdd
  *
  * 	out :
  * 	BroadcastResponse => data about broadcast
  *		error	
  */
func findBy(key string, val string) (BroadcastResponse, error) {
	var broadResponse BroadcastResponse

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return broadResponse, err
	}	

	err = db.QueryRow("SELECT id FROM broadcast WHERE " + key + " = ?", val).Scan(&broadResponse.ID)
	
	if err != nil {
		return broadResponse, err
	}

	return broadResponse, nil
}

type CreateBroadcastRequest struct {
	ID int64
	Name   string
    Description string
    Mails []string
}

/**
  * Create broadcast
  * 
  * 	in :
  *		w => ResponseWriter
  * 	createBroadcastRequest => params request
  */
func CreateBroadcast(w http.ResponseWriter, createBroadcastRequest CreateBroadcastRequest) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	res, err := db.Exec("INSERT `broadcast`(`name`, `description`) VALUES (?, ?)", createBroadcastRequest.Name, createBroadcastRequest.Description)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la liste de diffusion n'a pas été effectué")
		return 
	}

	broadcast_id, err := res.LastInsertId()
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, mail introuvable")
		return 
	}

	createBroadcastRequest.ID = broadcast_id

	for mailIndex := 0; mailIndex < len(createBroadcastRequest.Mails); mailIndex++ {
		subscriber, err := sub.FindBy("mail", createBroadcastRequest.Mails[mailIndex])
		if err != nil {
			service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, mail introuvable")
			return
		}

		insert, err := db.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broadcast_id, subscriber.ID)

		if err != nil {
			service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la création de la liste de diffusion n'a pas été effectué")
			return
		}

		defer insert.Close()
	}

	service.WriteJSON(w, http.StatusOK, createBroadcastRequest)
}

type SubRequest struct {
	BroadcastName string
	SubscriberMail string
}

/**
  * Add subscriber from broadcast
  * 
  * 	in :
  *		w => ResponseWriter
  * 	SubRequest => params request
  */
func AddSubscriber(w http.ResponseWriter, subRequest SubRequest) {
	subscriber, err := sub.FindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Subscriber introuvable")
		return
	}

	broad, err := findBy("name", subRequest.BroadcastName)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Liste de diffusion introuvable")
		return
	}

	db, err := config.ConnectToBDD()

	insert, err := db.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broad.ID, subscriber.ID)

	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, l'ajout d'un subscriber dans la liste n'a pas pu aboutir")
		return
	}

	service.WriteJSON(w, http.StatusOK, subRequest)

	defer insert.Close()
}

/**
  * Delete subscriber from broadcast
  * 
  * 	in :
  *		w => ResponseWriter
  * 	SubRequest => params request
  */
func DeleteSubscriber(w http.ResponseWriter, subRequest SubRequest) {
	subscriber, err := sub.FindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Subscriber introuvable")
		return
	}

	broad, err := findBy("name", subRequest.BroadcastName)
	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, Liste de diffusion introuvable")
		return
	}

	db, err := config.ConnectToBDD()

	delete, err := db.Query("DELETE FROM `broadcast_subscriber` WHERE broadcast_id = ? AND subscriber_id = ?", broad.ID, subscriber.ID)

	if err != nil {
		service.WriteErrorJSON(w, http.StatusInternalServerError, "Une erreur est survenue, la suppression d'un subscriber dans la liste n'a pas pu aboutir")
		return		
	}

	service.WriteJSON(w, http.StatusOK, subRequest)

	defer delete.Close()
}