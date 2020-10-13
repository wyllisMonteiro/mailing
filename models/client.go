package models

import (
	"log"
)

// Client : Structure format in db
type Client struct {
	ID       int    `json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// ClientFindBy : Get client according to params and return client
func ClientFindBy(key string, val string) (Client, error) {
	var client Client

	err := DB.QueryRow("SELECT id, mail, password FROM client WHERE "+key+" = ?", val).Scan(&client.ID, &client.Mail, &client.Password)

	if err != nil {
		log.Println(err)
		return client, err
	}

	return client, nil
}

// UpdateToken : update token client
func UpdateToken(token string, clientID int) error {
	update, err := DB.Query("UPDATE `client` SET `token` = ? WHERE `client`.`id` = ?", token, clientID)

	if err != nil {
		return err
	}

	defer update.Close()

	return nil
}
