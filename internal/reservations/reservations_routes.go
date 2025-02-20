package reservations

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type ReservationRoutes struct {
	router  *mux.Router
	handler *ReservationHandler
}

func NewRoutes(router *mux.Router, handler *ReservationHandler) *ReservationRoutes {
	return &ReservationRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *ReservationRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/reservation").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllReservation).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetReservationByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateReservation).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateReservation).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteReservation).Methods("DELETE")
}
