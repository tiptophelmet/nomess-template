package middleware

import (
	"net/http"
)

type mwFunc func(http.ResponseWriter, *http.Request)

var mw = map[string][]mwFunc{
	"/register": {WithLocalize},
	"/chat":     {WithSession},
}

func WithMiddleware(w http.ResponseWriter, r *http.Request) {
	mwList, found := mw[r.URL.Path]

	if !found {
		// default middleware
		mwList = []mwFunc{WithSession, WithLocalize}
	}

	for _, mw := range mwList {
		mw(w, r)
	}
}
