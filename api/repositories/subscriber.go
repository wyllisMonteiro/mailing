package repositories

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type Subscriber struct {
    ID   int    `json:"id"`
    Mail string `json:"mail"`
    Name string `json:"name"`
}

var subscriber Subscriber

func GetOneSubscriber(mail string) (Subscriber, error) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return subscriber, err
	}

	err = db.QueryRow("SELECT id, mail, name FROM subscriber WHERE mail = ?", mail).Scan(&subscriber.ID, &subscriber.Mail, &subscriber.Name)
	
	if err != nil {
		return subscriber, err
	}

	return subscriber, nil
}