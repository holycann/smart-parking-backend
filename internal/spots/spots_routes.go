package spots

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/spot", r.handler.HandleGetAllSpot).Methods("GET")
	r.router.HandleFunc("/spot/{id:[0-9]+}", r.handler.HandleGetSpotByID).Methods("GET")
	r.router.HandleFunc("/spot", r.handler.HandleCreateSpot).Methods("POST")
	r.router.HandleFunc("/spot/{id:[0-9]+}", r.handler.HandleUpdateSpot).Methods("PUT")
	r.router.HandleFunc("/spot/{id:[0-9]+}", r.handler.HandleDeleteSpot).Methods("DELETE")
}
