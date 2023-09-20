package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v5/intl"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/responder"
	"github.com/tiptophelmet/nomess-template/body"
	"github.com/tiptophelmet/nomess-template/response"
	"github.com/tiptophelmet/nomess-template/service"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Attempt to create an item...")

	var createItemBody body.CreateItem

	if err := json.NewDecoder(r.Body).Decode(&createItemBody); err != nil {
		logger.Debug("Invalid request body: %s", err.Error())

		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(response.BadRequest(), http.StatusBadRequest)
		return
	}

	srv := service.InitCreateItemService()

	err := srv.Validate(createItemBody)
	if err != nil {
		logger.Debug("Fields for creating item are invalid: %s", err.Error())

		badRequestResp := response.BadRequest()
		badRequestResp.Message = intl.Localize("invalid_body.message")
		badRequestResp.Text = err.Error()

		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(badRequestResp, http.StatusBadRequest)
		return
	}

	item, err := srv.Create(createItemBody)

	if err != nil {
		logger.Error("Could not create an item: %s", err.Error())

		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(response.InternalServerError(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Item created!")

	w.Header().Add("Content-Type", "application/json")
	responder.Respond(w, r)(response.ItemCreated(item), http.StatusOK)
}
