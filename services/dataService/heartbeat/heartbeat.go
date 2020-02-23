package heartbeat

import (
	"os"
	"time"

	"tonyxiong.top/gostorage/pkg/rabbitmq"
)

// StartHeartbeat start heartbeat
func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()

	for {
		q.Publish(
			"apiServers",                // exchange
			os.Getenv("LISTEN_ADDRESS"), // body
		)
		time.Sleep(5 * time.Second)
	}
}
