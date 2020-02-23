package object

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"tonyxiong.top/gostorage/pkg/stream"
	"tonyxiong.top/gostorage/services/apiService/heartbeat"
)

// Handler handl api request
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}

	if m == http.MethodGet {
		get(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	status, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(status)
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}

func putStream(object string) (*stream.PutStream, error) {
	server := heartbeat.ChooseRandomServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return stream.NewPutStream(server, object), nil
}
