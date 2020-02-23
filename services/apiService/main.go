package main

import (
	"log"
	"net/http"
	"os"

	"tonyxiong.top/gostorage/services/apiService/heartbeat"
	"tonyxiong.top/gostorage/services/apiService/locate"
	"tonyxiong.top/gostorage/services/dataService/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
