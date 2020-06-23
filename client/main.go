package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
	//"io/ioutil"
	"strconv"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecret")
var serverUrl = "http://localhost:9000/"

type User struct {
    ID   int    `json:"id"`
    Login string `json:"login"`
    Password string `json:"password"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	/*client := &http.Client{}
	req, _ := http.NewRequest("GET", serverUrl, nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}*/

	fmt.Fprintf(w, validToken)
	//fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Wyllis Monteiro"
	claims["pass"] = "mdp"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func connectToBDD() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mailing")

    if err != nil {
		return nil, err
	} else {
		fmt.Println("Logged to database")
		return db, nil
	}
}

func getOneUser(login string, password string) (User, error){

	db, err := connectToBDD()
	
	defer db.Close()

	var user User

	if err != nil {
		user = User {}
		return user, err
	}

	err = db.QueryRow("SELECT * FROM user WHERE login = ? AND password = ?", login, password).Scan(&user.ID, &user.Login, &user.Password)
	
	if err != nil {
		user = User {}
		return user, err
	}

	return user, nil 
}

func main() {
	fmt.Println("My Simple Client")

	user, err := getOneUser("wyllis", "w")

	if err != nil {
		panic(err.Error())
		return
	}

	log.Printf(strconv.Itoa(user.ID))
	log.Printf(user.Login)
	log.Printf(user.Password)
	
	handleRequests()
}