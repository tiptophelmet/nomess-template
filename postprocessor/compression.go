package postprocessor

import "net/http"

func WithCompression(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Add("Content-Encoding", "gzip, deflate")

	return w, r
}
