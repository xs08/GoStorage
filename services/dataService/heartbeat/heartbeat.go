package heartbeat

import (
	"os"
	"time"

	"tonyxiong.top/gostorage/pkg/rabbitmq"
)

// StartHeartbeat start heartbeat
func StartHeartbeat() {
	amqpLink := "amqp://admin:admin@" +
		os.Getenv("RABBITMQ_PORT_5672_TCP_ADDR") +
		":" + os.Getenv("RABBITMQ_PORT_5672_TCP_PORT")

	q := rabbitmq.New(amqpLink)
	defer q.Close()

	for {
		q.Publish(
			"apiServers",                // exchange
			os.Getenv("LISTEN_ADDRESS"), // body
		)
		time.Sleep(5 * time.Second)
	}
}
