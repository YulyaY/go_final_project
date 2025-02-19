package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, valueJson)
	p := password{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if p.Password != h.appCfg.AppPassword {
		w.WriteHeader(http.StatusUnauthorized)
		respBytes := responseErrorWrapper{ErrMsg: errInvalidPassword.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	token, err := jwtToken.SignedString([]byte(appConfig.Secret))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	res := tokenRequest{Token: token}
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
