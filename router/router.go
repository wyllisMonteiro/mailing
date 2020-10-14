package router

import (
	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/mailing/controllers"
)

// InitRoutes : Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/campaigns", controllers.CreateCampaign).Methods("POST")
	r.HandleFunc("/campaigns/{id}", controllers.GetCampaign).Methods("GET")
	r.HandleFunc("/broadcasts", controllers.CreateBroadcast).Methods("POST")
	r.HandleFunc("/broadcasts", controllers.GetBroadcast).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/broadcasts/subscriber", controllers.AddSubscriber).Methods("POST")
	r.HandleFunc("/broadcasts/subscriber", controllers.DeleteSubscriber).Methods("DELETE")

	return r
}
