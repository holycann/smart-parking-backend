package vehicles

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/vehicle", r.handler.HandleGetAllVehicle).Methods("GET")
	r.router.HandleFunc("/vehicle/{id:[0-9]+}", r.handler.HandleGetVehicleByID).Methods("GET")
	r.router.HandleFunc("/vehicle", r.handler.HandleCreateVehicle).Methods("POST")
	r.router.HandleFunc("/vehicle/{id:[0-9]+}", r.handler.HandleUpdateVehicle).Methods("PUT")
	r.router.HandleFunc("/vehicle/{id:[0-9]+}", r.handler.HandleDeleteVehicle).Methods("DELETE")
}
