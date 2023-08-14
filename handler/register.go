package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess/body"
	"github.com/tiptophelmet/nomess/intl"
	"github.com/tiptophelmet/nomess/responder"
	"github.com/tiptophelmet/nomess/response"
	"github.com/tiptophelmet/nomess/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerBody body.Register

	if json.NewDecoder(r.Body).Decode(&registerBody) != nil {
		responder.Respond(response.BadRequestResponse(), http.StatusBadRequest)
		return
	}

	srv := service.InitRegisterService()

	err := srv.Validate(registerBody)
	if err != nil {
		badRequestResp := response.BadRequestResponse()
		badRequestResp.Message = intl.Localize("invalid_body.message")
		badRequestResp.Text = err.Error()

		responder.Respond(badRequestResp, http.StatusBadRequest)
	}

	if srv.Register(registerBody) != nil {
		responder.Respond(response.InternalServerErrorResponse(), http.StatusInternalServerError)
	}

	responder.Respond(response.RegistrationSuccessfulResponse(), http.StatusOK)
}
