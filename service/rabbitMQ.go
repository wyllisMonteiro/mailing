package service

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
	"github.com/wyllisMonteiro/mailing/models"
)

// FailOnError : show error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// SendIDCampaign : Send ID campaign to RabbitMQ
func SendIDCampaign(idCampaign int64) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
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

// ReceiveMails : Consumer wait to get campaign id to send mails
func ReceiveMails() {

	time.Sleep(15 * time.Second)

	fmt.Println("Starting...")

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Println(err)
	}

	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			camp, err := models.CampaignFindByID(1)
			if err != nil {
				log.Printf("ERROR GET CAMPAIGN")
			}

			log.Printf("Received a message: %s", camp.Message)

			log.Printf("Received a message: %s", d.Body)

			subject := "send mail"
			body := "Hello <b>Bob</b> and <i>Cora</i>!"
			err = SendMail(subject, body)
			if err != nil {
				log.Println(err.Error())
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
