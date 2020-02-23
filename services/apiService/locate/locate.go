package locate

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"tonyxiong.top/gostorage/pkg/rabbitmq"
)

// Handler api request
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}

// Locate object data servers
func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	msg := <-c
	s, _ := strconv.Unlock(string(msg.Body))
	return s
}

// Exists of object data
func Exists(name string) bool {
	return Locate(name) != ""
}
