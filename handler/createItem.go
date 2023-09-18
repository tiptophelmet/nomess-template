package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v3/intl"
	"github.com/tiptophelmet/nomess-core/v3/responder"
	"github.com/tiptophelmet/nomess-template/body"
	"github.com/tiptophelmet/nomess-template/response"
	"github.com/tiptophelmet/nomess-template/service"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var createItemBody body.CreateItem

	if json.NewDecoder(r.Body).Decode(&createItemBody) != nil {
		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(response.BadRequest(), http.StatusBadRequest)
		return
	}

	srv := service.InitCreateItemService()

	err := srv.Validate(createItemBody)
	if err != nil {
		badRequestResp := response.BadRequest()
		badRequestResp.Message = intl.Localize("invalid_body.message")
		badRequestResp.Text = err.Error()

		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(badRequestResp, http.StatusBadRequest)
		return
	}

	item, err := srv.Create(createItemBody)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(response.InternalServerError(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	responder.Respond(w, r)(response.ItemCreated(item), http.StatusOK)

}
