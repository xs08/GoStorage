package stream

import (
	"fmt"
	"io"
	"net/http"
)

// PutStream put stream
type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

// NewPutStream get a put stream
func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)

	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		client := &http.Client{}
		r, e := client.Do(request)

		if e != nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)
		}
		c <- e
	}()

	return &PutStream{
		writer: writer,
		c:      c,
	}
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

// Close close put stream
func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}
