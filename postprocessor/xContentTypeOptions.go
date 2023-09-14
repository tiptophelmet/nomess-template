package postprocessor

import (
	"net/http"
)

func WithXContentTypeOptions(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Add("X-Content-Type-Options", "nosniff")

	return w, r
}
