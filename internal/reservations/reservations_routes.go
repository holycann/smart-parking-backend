package reservations

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/reservation", r.handler.HandleGetAllReservation).Methods("GET")
	r.router.HandleFunc("/reservation/{id:[0-9]+}", r.handler.HandleGetReservationByID).Methods("GET")
	r.router.HandleFunc("/reservation", r.handler.HandleCreateReservation).Methods("POST")
	r.router.HandleFunc("/reservation/{id:[0-9]+}", r.handler.HandleUpdateReservation).Methods("PUT")
	r.router.HandleFunc("/reservation/{id:[0-9]+}", r.handler.HandleDeleteReservation).Methods("DELETE")
}
