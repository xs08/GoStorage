package objects

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// Handler objects request
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
	logger.Info().Msg("request method not allowed")
}

func put(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Split(r.URL.EscapedPath(), "/")[2]
	filePath := os.Getenv("STORAGE_ROOT") + "/objects/" + fileName
	// logger
	log := logger.Info().
		Str("fileName", fileName).
		Str("filePath", filePath)

	f, e := os.Create(filePath)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Err(e).Msg("create file file error")
		return
	}
	defer f.Close()

	io.Copy(f, r.Body)
	w.WriteHeader(http.StatusOK)
	log.Msg("put file ok")
}

func get(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Split(r.URL.EscapedPath(), "/")[2]
	filePath := os.Getenv("STORAGE_ROOT") + "/objects/" + fileName
	// logger
	log := logger.Info().
		Str("fileName", fileName).
		Str("filePath", filePath)

	// file exists
	_, e := os.Stat(filePath)
	if e != nil {
		if os.IsNotExist(e) {
			w.WriteHeader(http.StatusNotFound)
			log.Msg("get file fail: not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Err(e).Msg("get file error")
		return
	}

	f, e := os.Open(filePath)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Err(e).Msg("open file error")
		return
	}
	defer f.Close()

	io.Copy(w, f)
	log.Msg("get file ok")
}
