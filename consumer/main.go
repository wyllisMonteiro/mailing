package main

import (
	"github.com/wyllisMonteiro/mailing/models"
	"github.com/wyllisMonteiro/mailing/service"
)

// Start consumer for waiting campaign id to send mails
func main() {
	models.ConnectToBDD()
	service.ReceiveMails()
}
