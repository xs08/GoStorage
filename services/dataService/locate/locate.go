package locate

import (
	"os"
	"strconv"

	"tonyxiong.top/gostorage/pkg/rabbitmq"
)

// Locate resource
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// StartLocate start location service
func StartLocate() {
	amqpLink := "amqp://admin:admin@" +
		os.Getenv("RABBITMQ_PORT_5672_TCP_ADDR") +
		":" + os.Getenv("RABBITMQ_PORT_5672_TCP_PORT")

	q := rabbitmq.New(amqpLink)
	defer q.Close()

	q.Bind("dataServers")
	c := q.Consume()

	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
