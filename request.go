package veriff

import (
	"net/http"
)

// request is a wrapper around http.Request.
type request struct {
	req      *http.Request
	decodeTo interface{}
}

func NewRequest(r *http.Request) *request {
	return &request{req: r}
}

func (r *request) DecodeTo(to interface{}) {
	r.decodeTo = to
}
