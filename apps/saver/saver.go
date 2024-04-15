package main

import (
	"database/sql"
	"log"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
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

	// Conexión a RabbitMQ
	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" +rabbit_password + "@" + rabbit_host + ":" + rabbit_port +"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Canal de comunicación con RabbitMQ
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaración de la cola para recibir mensajes de app2
	q, err := ch.QueueDeclare(
		"Cola2",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// Conexión a la base de datos MySQL
	db, err := sql.Open("postgres", "postgres://yunior:1234@db:5432/rabbitmq_go?sslmode=disable")
	failOnError(err, "Failed to connect to DB")
	defer db.Close()

	// Crear tabla si no existe
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id SERIAL PRIMARY KEY, content TEXT NOT NULL, received_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	failOnError(err, "Failed to create table")

	// Consumir mensajes de la cola
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

	// Procesar mensajes recibidos
	for d := range msgs {
		receivedMessage := green(string(d.Body))
		message := yellow("Received a message: ") + receivedMessage
		log.Print(message)

		// Guardar el mensaje en la base de datos junto con la hora actual
		_, err := db.Exec("INSERT INTO messages (content) VALUES ($1)", string(d.Body))
		failOnError(err, "Failed to insert message into DB")

		log.Println(blue("Message saved successfully"))

		// Enviar un mensaje a app1 para informar que el mensaje fue guardado correctamente
		err = ch.Publish(
			"",
			"Cola3", // Nombre de la cola para enviar mensajes a app1
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Message saved successfully"),
			},
		)
		failOnError(err, "Failed to publish a message")
	}
}