package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fajrinajiseno/mygolangapp/internal/entity"
	authUsecase "github.com/fajrinajiseno/mygolangapp/internal/module/auth/usecase"
	"github.com/fajrinajiseno/mygolangapp/internal/openapigen"
	"github.com/fajrinajiseno/mygolangapp/internal/transport"
)

type AuthHandler struct {
	authUC authUsecase.AuthUsecase
}

func NewAuthHandler(authUC authUsecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
	}
}

func (a *AuthHandler) PostV1AuthRegister(w http.ResponseWriter, r *http.Request) {
	var req openapigen.PostV1AuthRegisterJSONBody
	if !decodeJSONBody(w, r, &req) {
		return
	}
	message, err := a.authUC.Register(req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(openapigen.RegisterResponse{Message: &message})
	if err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}

func (a *AuthHandler) PostV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	var req openapigen.PostV1AuthLoginJSONBody
	if !decodeJSONBody(w, r, &req) {
		return
	}
	token, user, err := a.authUC.Login(req.Email, req.Password)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(openapigen.LoginResponse{Id: &user.ID, Email: &user.Email, Token: &token})
	if err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst any) bool {
	if r.Body == nil {
		transport.WriteAppError(w, entity.ErrorBadRequest("empty body"))
		return false
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		transport.WriteAppError(w, entity.ErrorBadRequest("failed to read body"))
		return false
	}

	if err := json.Unmarshal(body, dst); err != nil {
		transport.WriteAppError(w, entity.ErrorBadRequest("invalid json: "+err.Error()))
		return false
	}
	return true
}
