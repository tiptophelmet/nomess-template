package postprocessor

import (
	"net/http"
)

type postProcFunc func(http.ResponseWriter, *http.Request)

var postProc = map[string][]postProcFunc{
	"/api/auth/register": {WithLogging},
}

func WithPostProcessor(w http.ResponseWriter, r *http.Request) {
	postProcList, found := postProc[r.URL.Path]

	if !found {
		// default postprocessor
		postProcList = []postProcFunc{WithLogging}
	}

	for _, proc := range postProcList {
		proc(w, r)
	}
}
