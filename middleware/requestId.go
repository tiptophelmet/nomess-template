package middleware

import (
	"net/http"

	"github.com/tiptophelmet/nomess-template/util"
)

func WithRequestID(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	r = util.IssueRequestID(r)
	return w, r
}
