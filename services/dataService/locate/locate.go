package locate

import (
	"os"
	"strconv"

	"tonyxiong.top/gostorage/pkg/netutils"
	"tonyxiong.top/gostorage/pkg/rabbitmq"
	"tonyxiong.top/gostorage/pkg/utils"
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

	// connect to mq
	q := rabbitmq.New(amqpLink)
	defer q.Close()
	log.Msg("connect to rabbitmq ok")

	// bing exchange
	q.Bind("dataServers")
	c := q.Consume()

	// logger
	log = logger.Info().
		Str("exchange", "dataServers")

	// get inter ip
	ipAddr, err := netutils.GetNetInterIPV4()
	if err != nil {
		log.Err(err).Msg("locate object fail: can't get local ip")
	}
	log.Str("ipv4", ipAddr).Msg("get ipv4")

	for msg := range c {
		// receive message
		msgBody := string(msg.Body)
		log = log.Str("rawData", msgBody)

		object, e := strconv.Unquote(msgBody)
		if e != nil {
			log.Err(e).Msg("unquote msg error")
			break
		}
		log = log.Str("object", object)

		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {

			// replay message
			message := utils.StringsJoinWithBuilder([]string{ipAddr, os.Getenv("LISTEN_ADDRESS")})
			q.Send(msg.ReplyTo, message)

			log.Str("replayMessage", message).Msg("locate object ok")
			continue
		} else {
			log.Msg("locate object fail")
			continue
		}
	}
}
