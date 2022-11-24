package repository

import (
	"log"
	"github.com/streadway/amqp"
)

func RabbitInit() (ch *amqp.Channel, q amqp.Queue) {
	conn, _ := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	q, err = ch.QueueDeclare(
		"messages", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	return ch, q
}