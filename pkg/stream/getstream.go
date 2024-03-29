package stream

import (
	"fmt"
	"io"
	"net/http"
)

// GetStream get stream
type GetStream struct {
	reader io.Reader
}

func newGetStream(url string) (*GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}

	return &GetStream{r.Body}, nil
}

// NewGetStream get new get stream
func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}

	return newGetStream("http://" + server + "/objects/" + object)
}

func (r *GetStream) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}
