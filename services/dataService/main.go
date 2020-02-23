package main

import (
	"log"
	"net/http"
	"os"

	"tonyxiong.top/gostorage/services/dataService/heartbeat"
	"tonyxiong.top/gostorage/services/dataService/locate"
	"tonyxiong.top/gostorage/services/dataService/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()

	http.HandleFunc("/objects", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
