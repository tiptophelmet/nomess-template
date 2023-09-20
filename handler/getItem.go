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

func GetItem(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Attempt to get an item...")

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

	srv := service.InitGetItemService()
	item, err := srv.Get(int(itemId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("Item was not found")

			responder.Respond(w, r)(response.ItemNotFound(), http.StatusOK)
			return
		}

		logger.Error("Failed to get an item: %s", err)

		responder.Respond(w, r)(response.InternalServerError(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Item found!")

	w.Header().Add("Content-Type", "application/json")
	responder.Respond(w, r)(item, http.StatusOK)
}
