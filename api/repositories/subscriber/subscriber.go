package subscriber

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type SubscriberResponse struct {
    ID   int    `json:"id"`
    Mail string `json:"mail"`
    Name string `json:"name"`
}

var subscriber SubscriberResponse

func FindBy(key string, val string) (SubscriberResponse, error) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return subscriber, err
	}

	err = db.QueryRow("SELECT id, mail, name FROM subscriber WHERE " + key + " = ?", val).Scan(&subscriber.ID, &subscriber.Mail, &subscriber.Name)
	
	if err != nil {
		return subscriber, err
	}

	return subscriber, nil
}