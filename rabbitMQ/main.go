package main

import (
	service "github.com/wyllisMonteiro/mailing/rabbitMQ/service"
)

func main() {
	service.ReceiveMails()
}