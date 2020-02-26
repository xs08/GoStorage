package main

import (
	"log"
	"net/http"
	"os"

	"tonyxiong.top/gostorage/services/dataService/heartbeat"
	"tonyxiong.top/gostorage/services/dataService/locate"
	"tonyxiong.top/gostorage/services/dataService/objects"

	"tonyxiong.top/gostorage/pkg/logs"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()

	// add log middleware
	logMiddlewares := logs.GetHTTPLoggerMiddleware(os.Stdout, map[string]string{
		"appName":    "dataService",
		"appAddress": os.Getenv("LISTEN_ADDRESS"),
	})

	http.Handle("/objects/", logMiddlewares.Then(http.HandlerFunc(objects.Handler)))

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
