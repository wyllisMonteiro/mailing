package router

import (
	"github.com/wyllisMonteiro/mailing/client/controllers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", controllers.HomePage)
}
