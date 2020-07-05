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

func ClientFindBy(key string, val string) (Client, error) {
	var client Client

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return client, err
	}

	err = db.QueryRow("SELECT id, mail, password FROM client WHERE " + key + " = ?", val).Scan(&client.ID, &client.Mail, &client.Password)
	
	if err != nil {
		return client, err
	}

	return client, nil
}

func UpdateToken(token string, client_id int) (error) {
	db, err := config.ConnectToBDD()

	defer db.Close()

	if err != nil {
		return err
	}

	update, err := db.Query("UPDATE `client` SET `token` = ? WHERE `client`.`id` = ?", token, client_id)

	if err != nil {
		return err
	}

	defer update.Close()

	return nil
}
