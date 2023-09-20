package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/responder"
	"github.com/tiptophelmet/nomess-template/response"
	"github.com/tiptophelmet/nomess-template/service"
	"gorm.io/gorm"
)

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Attempt to delete an item...")

	pathParams := mux.Vars(r)

	itemIdStr, found := pathParams["id"]
	if !found {
		logger.Debug("Item id was not found in path parameters")

		responder.Respond(w, r)(response.ItemIdAbsent(), http.StatusBadRequest)
		return
	}

	itemId, err := strconv.ParseInt(itemIdStr, 10, 0)
	if err != nil {
		logger.Debug("Item id is invalid, must be an integer")

		responder.Respond(w, r)(response.ItemIdInvalid(), http.StatusBadRequest)
		return
	}

	srv := service.InitDeleteItemService()
	err = srv.Delete(int(itemId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("Item was not found")

			responder.Respond(w, r)(response.ItemNotFound(), http.StatusOK)
			return
		}

		logger.Error("Failed to delete an item: %s", err)

		responder.Respond(w, r)(response.InternalServerError(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Item deleted!")

	w.Header().Add("Content-Type", "application/json")
	responder.Respond(w, r)(response.ItemDeleted(), http.StatusOK)
}
