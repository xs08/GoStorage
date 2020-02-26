package heartbeat

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"tonyxiong.top/gostorage/pkg/rabbitmq"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

// ListenHeartbeat listen heartbeat store api
func ListenHeartbeat() {
	amqpLink := "amqp://admin:admin@" +
		os.Getenv("RABBITMQ_PORT_5672_TCP_ADDR") +
		":" + os.Getenv("RABBITMQ_PORT_5672_TCP_PORT")

	logger.Debug().
		Str("amqpLink", amqpLink).
		Msg("combine amqpLink from env")

	q := rabbitmq.New(amqpLink)
	defer q.Close()

	q.Bind("apiServers")
	c := q.Consume()
	// start remove data
	go removeExpiredDataServer()

	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))

		if e != nil {
			panic(e)
		}
		logger.Debug().
			Str("dataServer", dataServer).
			Msg("consume data from mq")

		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}

}

// removeExpiredDataServer clear expired serverdata
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

// GetDataServers get data servers
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()

	ds := make([]string, 0)
	for s := range dataServers {
		ds = append(ds, s)
	}

	logger.Debug().
		Msg("GetDataServers with ds")

	return ds
}

// ChooseRandomServer choose random server
func ChooseRandomServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}

	return ds[rand.Intn(n)]
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
