package auth

import (
	"github.com/gorilla/mux"
)

type AuthRoutes struct {
	router  *mux.Router
	handler *AuthHandler
}

func NewRoutes(router *mux.Router, handler *AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *AuthRoutes) RegisterRoutes() {
	r.router.HandleFunc("/login", r.handler.HandleUserLogin).Methods("POST")
}
