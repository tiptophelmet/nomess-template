package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess/body"
	resp "github.com/tiptophelmet/nomess/response"
	"github.com/tiptophelmet/nomess/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerBody body.Register
	if json.NewDecoder(r.Body).Decode(&registerBody) != nil {
		resp.Error("bad_request", http.StatusBadRequest)
		return
	}
	srv := service.InitRegisterService()
	err := srv.Validate(registerBody)
	if err != nil {
		resp.Respond(resp.Response{Resp: "Invalid fields", Txt: err.Error()}, http.StatusBadRequest)
	}

	if srv.Register(registerBody) != nil {
		resp.Error("internal_server_error", http.StatusInternalServerError)
	}

	resp.Ok("registration_successful", http.StatusOK)
}
