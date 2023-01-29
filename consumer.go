package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	amqpConnection()

}

func amqpConnection() {

	fmt.Println("Setting RabbitMQ server")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()
	fmt.Println("Successfully connected to our RabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	consumeMessage(ch)

}

func consumeMessage(ch *amqp.Channel) {

	msgs, err := ch.Consume(
		"RabbitMQ Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("[*] - waiting for messages ... ")
	<-forever

}
