package main

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/api/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("My Simple Client")
	router.InitRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))
}