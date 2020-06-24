package controllers

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/api/service"
	"github.com/wyllisMonteiro/mailing/api/repositories"
	"net/http"
	"log"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user, err := repositories.GetOneUser("wyllis")
	if err != nil {
		panic(err.Error())
		return
	}

	match, err := service.ComparePasswordAndHash("w", user.Password)
    if err != nil {
		log.Fatal(err)
		return
	}

	if !match {
		return
	}

	validToken, err := service.GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	repositories.InsertUserToken(validToken, user.ID)

	/*client := &http.Client{}
	req, _ := http.NewRequest("GET", SERVER_URL, nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}*/

	_, _ = fmt.Fprintf(w, validToken)
	//fmt.Fprintf(w, string(body))
}