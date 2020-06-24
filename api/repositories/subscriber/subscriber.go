package subscriber

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type SubscriberResponse struct {
    ID   int    `json:"id"`
    Mail string `json:"mail"`
    Name string `json:"name"`
}

func FindBy(key string, val string) (SubscriberResponse, error) {
	var sub SubscriberResponse

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return sub, err
	}

	err = db.QueryRow("SELECT id, mail, name FROM subscriber WHERE " + key + " = ?", val).Scan(&sub.ID, &sub.Mail, &sub.Name)
	
	if err != nil {
		return sub, err
	}

	return sub, nil
}