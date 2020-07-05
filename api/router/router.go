package router

import (
	"log"
	"github.com/wyllisMonteiro/mailing/api/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/campaigns", controllers.CreateCampaign).Methods("POST")
	r.HandleFunc("/campaigns/{id}", controllers.GetCampaign).Methods("GET")
	r.HandleFunc("/broadcasts", controllers.CreateBroadcast).Methods("POST")
	r.HandleFunc("/broadcasts", controllers.GetBroadcast).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/broadcasts/subscriber", controllers.AddSubscriber).Methods("POST")
	r.HandleFunc("/broadcasts/subscriber", controllers.DeleteSubscriber).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
