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
	exchnage := "apiServers"
	// logger
	log := logger.Info().
		Str("amqpLink", amqpLink).
		Str("exchnage", exchnage)

	q := rabbitmq.New(amqpLink)
	defer q.Close()

	q.Bind(exchnage)
	c := q.Consume()
	log.Msg("start consume msg from mq")
	// start remove data
	go removeExpiredDataServer()

	// consume msg logger
	log = logger.Trace().
		Str("amqpLink", amqpLink).
		Str("exchnage", exchnage)
	for msg := range c {
		msg := string(msg.Body)
		log = log.Str("rawMsg", msg)

		dataServer, e := strconv.Unquote(msg)
		log = log.Str("dataServer", dataServer)

		if e != nil {
			log.Err(e).Msg("unquote msg error")
			panic(e)
		}

		log.Msg("consume data from mq")

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
	logger.Debug().Strs("dataServers", ds).Msg("get data servers with ds")

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
