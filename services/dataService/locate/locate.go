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
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
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
