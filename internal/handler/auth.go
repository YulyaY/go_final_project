package handler

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

const (
	envPassword = "TODO_PASSWORD"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv(envPassword)
		if len(pass) > 0 {
			var jwtT string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwtT = cookie.Value
			}

			jwtToken, err := jwt.Parse(jwtT, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if !jwtToken.Valid {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
