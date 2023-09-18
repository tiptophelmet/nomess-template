package handler

import (
	"net/http"

	"github.com/tiptophelmet/nomess-core/v4/responder"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	responder.Respond(w, r)("Hello World!", http.StatusOK)
}
