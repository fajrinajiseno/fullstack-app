package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fajrinajiseno/mygolangapp/internal/api"
	"github.com/fajrinajiseno/mygolangapp/internal/config"
	"github.com/fajrinajiseno/mygolangapp/internal/entity"
	ah "github.com/fajrinajiseno/mygolangapp/internal/module/auth/handler"
	aum "github.com/fajrinajiseno/mygolangapp/internal/module/auth/usecase/mock"
	srv "github.com/fajrinajiseno/mygolangapp/internal/service/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLoginAndAccessProtected(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const hour24 = 24
	claims := jwt.MapClaims{
		"sub":   "1",
		"email": "alice@example.com",
		"exp":   time.Now().Add(hour24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(config.JwtSecret)

	mockAuthUC := aum.NewMockAuthUsecase(ctrl)
	mockAuthUC.EXPECT().
		Login("alice@example.com", "password").
		Return(signed, &entity.User{
			ID:           1,
			Email:        "alice@example.com",
			PasswordHash: "password",
			CreatedAt:    time.Now(),
		}, nil)
	mockAuthUC.EXPECT().
		Register("alice@example.com", "password", "password").
		Return("Success Register", nil)

	authH := ah.NewAuthHandler(mockAuthUC)

	apiHandler := &api.APIHandler{
		Auth: authH,
	}

	srv := srv.NewServer(apiHandler, "../../../../openapi.yaml")
	ts := httptest.NewServer(srv.Routes())
	defer ts.Close()

	register := map[string]string{"email": "alice@example.com", "password": "password", "confirmPassword": "password"}
	registerJson, _ := json.Marshal(register)
	resRegister, err := http.Post(ts.URL+"/v1/auth/register", "application/json", bytes.NewReader(registerJson))
	require.NoError(t, err)
	defer resRegister.Body.Close()
	require.Equal(t, http.StatusOK, resRegister.StatusCode)

	var respRegister map[string]string
	json.NewDecoder(resRegister.Body).Decode(&respRegister)
	respMessage := respRegister["message"]
	require.NotEmpty(t, respMessage)

	login := map[string]string{"email": "alice@example.com", "password": "password"}
	loginJson, _ := json.Marshal(login)
	resLogin, err := http.Post(ts.URL+"/v1/auth/login", "application/json", bytes.NewReader(loginJson))
	require.NoError(t, err)
	defer resLogin.Body.Close()
	require.Equal(t, http.StatusOK, resLogin.StatusCode)

	var respLogin map[string]string
	json.NewDecoder(resLogin.Body).Decode(&respLogin)
	respToken := respLogin["token"]
	require.NotEmpty(t, respToken)

}
