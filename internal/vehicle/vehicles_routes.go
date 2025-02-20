package vehicles

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type VehicleRoutes struct {
	router  *mux.Router
	handler *VehicleHandler
}

func NewRoutes(router *mux.Router, handler *VehicleHandler) *VehicleRoutes {
	return &VehicleRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *VehicleRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/vehicle").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllVehicle).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetVehicleByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateVehicle).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateVehicle).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteVehicle).Methods("DELETE")
}
