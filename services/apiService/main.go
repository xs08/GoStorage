package main

import (
	"log"
	"net/http"
	"os"

	"tonyxiong.top/gostorage/services/apiService/heartbeat"
	"tonyxiong.top/gostorage/services/apiService/locate"
	"tonyxiong.top/gostorage/services/apiService/objects"

	"tonyxiong.top/gostorage/pkg/logs"
)

func main() {
	go heartbeat.ListenHeartbeat()

	// add log middleware
	logMiddlewares := logs.GetHTTPLoggerMiddleware(os.Stdout, map[string]string{
		"appName":    "apiService",
		"appAddress": os.Getenv("LISTEN_ADDRESS"),
	})

	http.Handle("/objects/", logMiddlewares.Then(http.HandlerFunc(objects.Handler)))
	http.Handle("/locate/", logMiddlewares.Then(http.HandlerFunc(locate.Handler)))

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
