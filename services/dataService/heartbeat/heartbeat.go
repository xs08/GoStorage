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
	// logger
	log := logger.Info().Str("amqpLink", amqpLink)

	q := rabbitmq.New(amqpLink)
	defer q.Close()
	log.Msg("connect to rabbitmq ok")

	for {
		q.Publish(
			"apiServers",                // exchange
			os.Getenv("LISTEN_ADDRESS"), // body
		)
		logger.Debug().Str("publishData", os.Getenv("LISTEN_ADDRESS")).Msg("send heart beat")
		time.Sleep(5 * time.Second)
	}
}
