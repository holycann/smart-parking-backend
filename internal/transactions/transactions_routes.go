package transactions

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type TransactionRoutes struct {
	router  *mux.Router
	handler *TransactionHandler
}

func NewRoutes(router *mux.Router, handler *TransactionHandler) *TransactionRoutes {
	return &TransactionRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *TransactionRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/transaction").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllTransaction).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetTransactionByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreateTransaction).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdateTransaction).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeleteTransaction).Methods("DELETE")
}
