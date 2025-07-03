package rabbit

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func NewQ() (*amqp091.Queue, *amqp091.Channel) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	queue, err := ch.QueueDeclare(
		"BET_SETTLEMENT", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		queue.Name,       // queue name
		"",               // routing key
		"BET_SETTLEMENT", // exchange
		false,
		nil,
	)

	return &queue, ch
}
