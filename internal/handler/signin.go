package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

const (
	secret = "0e3a308c7d4d4b4e48f6a1b29ca30ff0"
)

var errInvalidPassword error = errors.New("invalid password")

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := struct {
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	log.Printf("password: '%s', from env: '%s'", p.Password, os.Getenv(envPassword))
	if p.Password != os.Getenv(envPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		respBytes := responseErrorWrapper{ErrMsg: errInvalidPassword.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}
	result, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}
