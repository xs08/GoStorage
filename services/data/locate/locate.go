package locate

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
	"tonyxiong.top/gostorage/pkg/logs"
)

// Locate resource
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
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

}
