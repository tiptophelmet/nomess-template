package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiptophelmet/nomess-core/v5/intl"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/responder"
	"github.com/tiptophelmet/nomess-template/body"
	"github.com/tiptophelmet/nomess-template/response"
	"github.com/tiptophelmet/nomess-template/service"
	"gorm.io/gorm"
)

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Attempt to get an item...")

	pathParams := mux.Vars(r)
	itemIdStr, found := pathParams["id"]

	if !found {
		logger.Debug("Item id was not found in path parameters")

		responder.Respond(w, r)(response.ItemIdAbsent(), http.StatusBadRequest)
		return
	}

	itemId, err := strconv.ParseUint(itemIdStr, 10, 0)
	if err != nil {
		logger.Debug("Item id is invalid, must be an integer")

		responder.Respond(w, r)(response.ItemIdInvalid(), http.StatusBadRequest)
		return
	}

	var updateItemBody body.UpdateItem

	if err := json.NewDecoder(r.Body).Decode(&updateItemBody); err != nil {
		logger.Debug("Invalid request body: %s", err.Error())

		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(response.BadRequest(), http.StatusBadRequest)
		return
	}

	srv := service.InitUpdateItemService()
	err = srv.Validate(updateItemBody)

	if err != nil {
		logger.Debug("Fields for updating item are invalid: %s", err.Error())

		badRequestResp := response.BadRequest()
		badRequestResp.Message = intl.Localize("invalid_body.message")
		badRequestResp.Text = err.Error()

		w.Header().Add("Content-Type", "application/json")
		responder.Respond(w, r)(badRequestResp, http.StatusBadRequest)
		return
	}

	updated, err := srv.Update(int(itemId), updateItemBody)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("Item was not found")

			responder.Respond(w, r)(response.ItemNotFound(), http.StatusOK)
			return
		}

		logger.Error("Failed to update an item: %s", err)

		responder.Respond(w, r)(response.InternalServerError(), http.StatusInternalServerError)
		return

	}
	logger.Debug("Item updated!")

	w.Header().Add("Content-Type", "application/json")
	responder.Respond(w, r)(response.ItemUpdated(updated), http.StatusOK)
}
