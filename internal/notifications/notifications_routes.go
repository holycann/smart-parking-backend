package notifications

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type NotificationRoutes struct {
	router  *mux.Router
	handler *NotificationHandler
}

func NewRoutes(router *mux.Router, handler *NotificationHandler) *NotificationRoutes {
	return &NotificationRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *NotificationRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/notification").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllNotifications).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateNotification).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateNotification).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteNotification).Methods("DELETE")
}
