package repositories

func insertUserToken(token string, user_id int) {
	db, err := connectToBDD()

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("UPDATE `user` SET `token` = ? WHERE `user`.`id` = ?", token, user_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
