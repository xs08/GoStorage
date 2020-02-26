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
	// logger
	log := logger.Info().Str("amqpLink", amqpLink)

	q := rabbitmq.New(amqpLink)
	defer q.Close()
	log.Msg("connect to rabbitmq ok")

	q.Bind("dataServers")
	c := q.Consume()

	log = logger.Info().
		Str("exchange", "dataServers")

	for msg := range c {
		msgBody := string(msg.Body)
		log = log.Str("rawData", msgBody)

		object, e := strconv.Unquote(msgBody)
		if e != nil {
			log.Err(e).Msg("unquote msg error")
			break
		}
		log = log.Str("object", object)

		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
			log.Msg("locate object ok")
			break
		} else {
			log.Msg("locate object fail")
			break
		}
	}
}
