package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/mailing/models"
	"github.com/wyllisMonteiro/mailing/router"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)

	models.ConnectToBDD()

	log.Fatal(http.ListenAndServe(":9000", r))
}
