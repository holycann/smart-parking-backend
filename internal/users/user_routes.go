package users

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/user", r.handler.HandleGetAllUsers).Methods("GET")
	r.router.HandleFunc("/user/{id:[0-9]+}", r.handler.HandleGetUserByID).Methods("GET")
	r.router.HandleFunc("/user", r.handler.HandleCreateUser).Methods("POST")
	r.router.HandleFunc("/user/{id:[0-9]+}", r.handler.HandleUpdateUser).Methods("PUT")
	r.router.HandleFunc("/user/{id:[0-9]+}", r.handler.HandleDeleteUser).Methods("DELETE")
}
