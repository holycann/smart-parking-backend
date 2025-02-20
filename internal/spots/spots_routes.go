package spots

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type SpotRoutes struct {
	router  *mux.Router
	handler *SpotHandler
}

func NewRoutes(router *mux.Router, handler *SpotHandler) *SpotRoutes {
	return &SpotRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *SpotRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/spot").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllSpot).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetSpotByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateSpot).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateSpot).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteSpot).Methods("DELETE")
}
