package service

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// FailOnError : show error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// SendIDCampaign : Send ID campaign to RabbitMQ
func SendIDCampaign(idCampaign int64) {
	fmt.Println(5555)
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	fmt.Println(6666)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // pass
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	body := "1"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
}
