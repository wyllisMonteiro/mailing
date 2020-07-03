package router

import (
	"log"
	"github.com/wyllisMonteiro/mailing/api/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Login).Methods("POST")
	r.HandleFunc("/campaign", controllers.Campaign).Methods("GET", "POST")
	r.HandleFunc("/broadcast", controllers.BroadCast).Methods("GET, POST")
	r.HandleFunc("/broadcast/add/subscriber", controllers.AddSubscriber).Methods("POST")
	r.HandleFunc("/broadcast/delete/subscriber", controllers.DeleteSubscriber).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
