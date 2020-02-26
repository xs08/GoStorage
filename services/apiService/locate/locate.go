package locate

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"tonyxiong.top/gostorage/pkg/rabbitmq"
)

// Handler api request
func Handler(w http.ResponseWriter, r *http.Request) {
	// logger
	log := logger.Info()

	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Msg("bad request for get object")
		return
	}

	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	log = log.Str("locateObject", info)

	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		log.Msg("object not found")
		return
	}

	b, _ := json.Marshal(info)
	w.Write(b)
	log.Str("jsonObject", string(b)).Msg("get Object ok")
}

// Locate object data servers
func Locate(name string) string {
	amqpLink := "amqp://admin:admin@" +
		os.Getenv("RABBITMQ_PORT_5672_TCP_ADDR") +
		":" + os.Getenv("RABBITMQ_PORT_5672_TCP_PORT")
	// logger
	log := logger.Info().
		Str("amqpLink", amqpLink).
		Str("locateFileName", name)

	q := rabbitmq.New(amqpLink)
	q.Publish("dataServers", name)
	c := q.Consume()

	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	msg := string((<-c).Body)
	log.Str("rawMsg", msg)

	s, _ := strconv.Unquote(msg)
	log.Str("unquteMsg", s).Msg("locate object")

	return s
}

// Exists of object data
func Exists(name string) bool {
	return Locate(name) != ""
}
