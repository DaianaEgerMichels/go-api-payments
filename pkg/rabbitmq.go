package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// open channel with rabbitmq
func OpenChannel() (*amqp.Channel, error) {
	// for case real, this information should be in environment variables
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")


	if err != nil {
		// command for stop the program
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

// consume messages
func Consume(ch *amqp.Channel, out chan amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue, // queue consumed name
		"go-payment",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}

// publish messages
func Publish(ctx context.Context, ch *amqp.Channel, body, exName string) error {
	err := ch.PublishWithContext(
		ctx,
		exName,
		"PaymentDone", // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body), // the body comes as a string, but must be sent in bytes 
		},
	)
	if err != nil {
		return err
	}
	return nil
}
