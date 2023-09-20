package handler

import (
	"errors"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/responder"
	"github.com/tiptophelmet/nomess-template/response"
	"github.com/tiptophelmet/nomess-template/service"
	"gorm.io/gorm"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	srv := service.InitGetItemsService()

	items, err := srv.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Debug("Items were not found")

			responder.Respond(w, r)(response.ItemsNotFound(), http.StatusOK)
			return
		}

		logger.Error("Failed to get items: %s", err)

		responder.Respond(w, r)(response.InternalServerError(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Items found!")

	w.Header().Add("Content-Type", "application/json")
	responder.Respond(w, r)(items, http.StatusOK)
}
