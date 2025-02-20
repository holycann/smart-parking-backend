package zones

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type ZoneRoutes struct {
	router  *mux.Router
	handler *ZoneHandler
}

func NewRoutes(router *mux.Router, handler *ZoneHandler) *ZoneRoutes {
	return &ZoneRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *ZoneRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/zone").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllZone).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetZoneByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateZone).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateZone).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteZone).Methods("DELETE")
}
