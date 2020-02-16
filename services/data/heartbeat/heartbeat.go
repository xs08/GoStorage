package heartbeat

import (
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"

	"tonyxiong.top/gostorage/pkg/logs"
)

// StartHeartbeat start heartbeat
func StartHeartbeat() {
	rabbitLink := "amqp://admin:admin@" + os.Getenv("RABBITMQ_PORT_15672_TCP_ADDR") +
		":" + os.Getenv("RABBITMQ_PORT_5672_TCP_PORT") + "/"
	fmt.Println("rabbitLink: " + rabbitLink)

	conn, err := amqp.Dial(rabbitLink)
	logs.FailExit(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	logs.FailExit(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	logs.FailExit(err, "Failed to declare a queue")

	msgData := os.Getenv("LISTEN_ADDRESS")
	msgBody := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgData),
	}
	fmt.Printf("msgBody: %s\n", msgBody)

	for {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			msgBody,
		)
		logs.FailError(err, "Publish msg error")
		time.Sleep(5 * time.Second)
	}
}
