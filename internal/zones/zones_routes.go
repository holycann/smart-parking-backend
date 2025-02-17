package zones

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/zone", r.handler.HandleGetAllZone).Methods("GET")
	r.router.HandleFunc("/zone/{id:[0-9]+}", r.handler.HandleGetZoneByID).Methods("GET")
	r.router.HandleFunc("/zone", r.handler.HandleCreateZone).Methods("POST")
	r.router.HandleFunc("/zone/{id:[0-9]+}", r.handler.HandleUpdateZone).Methods("PUT")
	r.router.HandleFunc("/zone/{id:[0-9]+}", r.handler.HandleDeleteZone).Methods("DELETE")
}
