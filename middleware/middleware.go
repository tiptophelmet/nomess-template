package middleware

import (
	"net/http"
)

type mwFunc func(*http.ResponseWriter, *http.Request)

var mw = map[string][]mwFunc{
	"/api/auth/register": {WithAuth, WithLocalize},
}

func WithMiddleware(w *http.ResponseWriter, r *http.Request, handlr func(http.ResponseWriter, *http.Request)) {
	mwList, found := mw[r.URL.Path]

	if !found {
		// default middleware
		mwList = []mwFunc{WithAuth, WithLocalize}
	}

	for _, mw := range mwList {
		mw(w, r)
	}

	handlr(*w, r)
}
