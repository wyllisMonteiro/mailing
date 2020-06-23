package controllers

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/client/service"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := service.GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	insertUserToken(validToken, user.ID)

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