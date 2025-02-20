package users

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type UserRoutes struct {
	router  *mux.Router
	handler *UserHandler
}

func NewRoutes(router *mux.Router, handler *UserHandler) *UserRoutes {
	return &UserRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *UserRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/user").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllUsers).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetUserByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateUser).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateUser).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteUser).Methods("DELETE")
}
