package client

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type ClientResponse struct {
    ID   int    `json:"id"`
    Mail string `json:"mail"`
    Password string `json:"password"`
    Token string `json:"token"`
}

var clientResponse ClientResponse

func FindBy(key string, val string) (ClientResponse, error) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return clientResponse, err
	}

	err = db.QueryRow("SELECT id, mail, password FROM client WHERE " + key + " = ?", val).Scan(&clientResponse.ID, &clientResponse.Mail, &clientResponse.Password)
	
	if err != nil {
		return clientResponse, err
	}

	return clientResponse, nil
}

func UpdateToken(token string, client_id int) {
	db, err := config.ConnectToBDD()

	defer db.Close()

	if err != nil {
		return
	}

	// perform a db.Query insert
	update, err := db.Query("UPDATE `client` SET `token` = ? WHERE `client`.`id` = ?", token, client_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer update.Close()
}
