package veriff

import (
	"io"
	"net/http"
)

// request is a wrapper around http.Request.
type request struct {
	req      *http.Request
	decodeTo interface{}
	pipeTo   io.Writer
}

func NewRequest(r *http.Request) *request {
	return &request{req: r}
}

func (r *request) DecodeTo(to interface{}) {
	r.decodeTo = to
}

func (r *request) PipeTo(to io.Writer) {
	r.pipeTo = to
}
