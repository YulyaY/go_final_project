package handler

import (
	"net/http"
	"github.com/YulyaY/go_final_project.git/internal/app"
	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/YulyaY/go_final_project.git/internal/domain/pkg"
)

func BuildAuthMiddleware(appConfig config.AppConfig, appSettings app.AppSettings) func(http.Handler) http.Handler {
	middlewareFunc := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if appSettings.IsAuthentificationControlSwitchedOn {
				var jwtT string
				cookie, err := r.Cookie(valueToken)
				if err == nil {
					jwtT = cookie.Value
					//tests.Token = jwtT
				}

				errOfValidJwtToken := pkg.CreateJwtToken(jwtT, appConfig.Secret)
				if errOfValidJwtToken != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}

	return middlewareFunc
}
