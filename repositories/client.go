package repositories

import (
	"github.com/wyllisMonteiro/mailing/config"
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

	db, err := config.ConnectToBDD()
	if err != nil {
		return client, err
	}

	defer db.Close()

	err = db.QueryRow("SELECT id, mail, password FROM client WHERE "+key+" = ?", val).Scan(&client.ID, &client.Mail, &client.Password)

	if err != nil {
		return client, err
	}

	return client, nil
}

// UpdateToken : update token client
func UpdateToken(token string, client_id int) error {
	db, err := config.ConnectToBDD()
	if err != nil {
		return err
	}

	defer db.Close()

	update, err := db.Query("UPDATE `client` SET `token` = ? WHERE `client`.`id` = ?", token, client_id)

	if err != nil {
		return err
	}

	defer update.Close()

	return nil
}
