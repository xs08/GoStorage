package heartbeat

import (
	"os"
	"time"

	"tonyxiong.top/gostorage/pkg/netutils"
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

	// get inter ip
	ipAddr, err := netutils.GetNetInterIPV4()
	if err != nil {
		log.Err(err).Msg("locate object fail: can't get local ip")
	}
	// listen address
	listenAddress := ipAddr + os.Getenv("LISTEN_ADDRESS")
	log.Str("ipv4", ipAddr).Str("address", listenAddress).Msg("get ipv4")

	for {

		q.Publish(
			"apiServers",  // exchange
			listenAddress, // body
		)
		logger.Trace().Str("publishData", os.Getenv("LISTEN_ADDRESS")).Msg("send heart beat")
		time.Sleep(5 * time.Second)
	}
}
