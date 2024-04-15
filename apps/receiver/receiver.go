package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/streadway/amqp"

	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT") 
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {

	//red := color.New(color.FgRed).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()
	blue := color.New(color.FgBlue).SprintfFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()



	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" +rabbit_password + "@" + rabbit_host + ":" + rabbit_port +"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Cola1",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

		log.Print(blue(" ||  Waiting for messages. To exit press CTRL+C"))

	for d := range msgs {
		receivedMessage := green(string(d.Body))
		message := yellow("Received a message: ") + receivedMessage
		log.Print(message)

		// Enviar el mensaje a app3
		err := ch.Publish(
			"",
			"Cola2", // Nombre de la cola para enviar mensajes a app3
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        d.Body,
			},
		)
		failOnError(err, "Failed to publish a message")
	}
}