package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"tonyxiong.top/gostorage/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Printf("Server Listen At: %s\n", os.Getenv("LISTEN_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
