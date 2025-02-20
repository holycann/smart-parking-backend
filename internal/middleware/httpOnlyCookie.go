package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/holycann/smart-parking-backend/config"
)

func SetJWTHttpOnlyCookie(w http.ResponseWriter, r *http.Request, token string, isUseHttps bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Second * time.Duration(config.Env.JWTExpirationInSecond)),
		HttpOnly: true,
		Path:     "/",
		Secure:   isUseHttps,
		SameSite: http.SameSiteStrictMode,
	})
}

func GetJWTHttpOnlyCookie(w http.ResponseWriter, r *http.Request, token string) (*http.Cookie, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		// http.Error(w, "", http.StatusUnauthorized)
		return nil, fmt.Errorf("Get JWT From Cookie Unauthorized!")
	}

	return cookie, nil
}
