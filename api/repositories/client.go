package repositories

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type Client struct {
    ID   int    `json:"id"`
    Mail string `json:"mail"`
    Password string `json:"password"`
    Token string `json:"token"`
}

var client Client

func GetOneClient(mail string) (Client, error) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return client, err
	}

	err = db.QueryRow("SELECT id, mail, password FROM client WHERE mail = ?", mail).Scan(&client.ID, &client.Mail, &client.Password)
	
	if err != nil {
		return client, err
	}

	return client, nil
}

func InsertClientToken(token string, client_id int) {
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
