package main

import (
	"fmt"
	"math/rand"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully Connected to out RabbitMQ Instance")

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plan",
			Body:        []byte(fmt.Sprintf("Hello %d", rand.Intn(100))),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected to out RabbitMQ to Queue")
}
