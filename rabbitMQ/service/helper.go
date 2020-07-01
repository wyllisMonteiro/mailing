package service

import (
	"log"
	"github.com/streadway/amqp"
	repos "github.com/wyllisMonteiro/mailing/api/repositories"
)

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func ReceiveMails() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
			campaignId :=d.Body
			camp, err = repos.FindByID(campaignId)
			if err != nil {
				log.Printf("ERROR GET CAMPAIGN")
			}

			log.Printf(camp.Message)
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}