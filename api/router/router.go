package router

import (
	"github.com/wyllisMonteiro/mailing/api/controllers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", controllers.Login)
	http.HandleFunc("/broadcast_list", controllers.BroadCastList)
}
