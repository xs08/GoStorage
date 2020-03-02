package objects

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"tonyxiong.top/gostorage/pkg/stream"
	"tonyxiong.top/gostorage/services/apiService/heartbeat"
	"tonyxiong.top/gostorage/services/apiService/locate"
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

// put object
func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	// logger
	log := logger.Info().Str("objectName", object)

	status, e := storeObject(r.Body, object)
	log = log.Int("status", status)

	if e != nil {
		log.Err(e).Msg("store object fail")
	}

	w.WriteHeader(status)
	log.Msg("store object finish")
}

// storeObject store object
func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	// logger
	log := logger.Info().Str("objectName", object)

	if e != nil {
		log.Err(e).Msg("create put stream fail")
		return http.StatusServiceUnavailable, e
	}

	io.Copy(stream, r)
	e = stream.Close()

	if e != nil {
		log.Err(e).Msg("close put stream fail")
		return http.StatusInternalServerError, e
	}

	log.Msg("store object finish")
	return http.StatusOK, nil
}

func putStream(object string) (*stream.PutStream, error) {
	server := heartbeat.()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	logger.Info().Str("RandomServer", server).Msg("get random server")

	return stream.NewPutStream(server, object), nil
}

// get objects
func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	// logger
	log := logger.Info().Str("objectName", object)

	stream, e := getStream(object)
	if e != nil {
		log.Err(e).Msg("locate object fail")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	io.Copy(w, stream)
	log.Msg("locate object success")
}

// create getStream
func getStream(object string) (*stream.GetStream, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}

	return stream.NewGetStream(server, object)
}
