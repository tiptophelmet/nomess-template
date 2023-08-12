package response

import (
	"encoding/json"
	"net/http"
)

type Responder struct {
	w *http.ResponseWriter
}

var resp *Responder

func initResponder(w *http.ResponseWriter) *Responder {
	if resp != nil {
		return resp
	}

	resp = &Responder{w}
	return resp
}

func Respond(r Response, statusCode int) {
	(*resp.w).WriteHeader(statusCode)
	json.NewEncoder(*resp.w).Encode(r)
}
