package main

import (
	"github.com/wyllisMonteiro/mailing/models"
	"github.com/wyllisMonteiro/mailing/service"
)

func main() {
	models.ConnectToBDD()
	service.ReceiveMails()
}
