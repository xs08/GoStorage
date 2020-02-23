package heartbeat

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

// ListenHeartbeat listen heartbeat store api
func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
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
