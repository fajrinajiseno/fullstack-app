package api

import (
	"net/http"

	ah "github.com/fajrinajiseno/mygolangapp/internal/module/auth/handler"
	"github.com/fajrinajiseno/mygolangapp/internal/openapigen"
)

type APIHandler struct {
	Auth *ah.AuthHandler
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostV1AuthRegister(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostV1AuthRegister(w, r)
}

func (h *APIHandler) PostV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostV1AuthLogin(w, r)
}
