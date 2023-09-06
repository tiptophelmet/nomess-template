package responder

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess/internal/postprocessor"
)

type Responder struct {
	w http.ResponseWriter
	r *http.Request
}

var resp *Responder

func Init(w http.ResponseWriter, r *http.Request) *Responder {
	if resp != nil {
		return resp
	}

	resp = &Responder{w, r}
	return resp
}

func Respond(r interface{}, statusCode int) {
	postprocessor.WithPostProcessor(resp.w, resp.r)

	resp.w.WriteHeader(statusCode)
	json.NewEncoder(resp.w).Encode(r)
}
