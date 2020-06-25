package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/wyllisMonteiro/mailing/api/service"
	sub "github.com/wyllisMonteiro/mailing/api/repositories/subscriber"
	"github.com/wyllisMonteiro/mailing/api/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("My Simple Client")
	//rabbitSend()

	router.InitRoutes()
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func rabbitSend() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	service.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	service.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // pass
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	  service.FailOnError(err, "Failed to declare a queue")
	  
	  test, err := sub.FindBy("name", "Kevin")

	  if err != nil {
		fmt.Println(test.Name)
		return
	  }

	  body := test.Name
	  err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        []byte(body),
		})
		service.FailOnError(err, "Failed to publish a message")
}