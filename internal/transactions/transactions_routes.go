package transactions

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/transaction", r.handler.HandleGetAllTransaction).Methods("GET")
	r.router.HandleFunc("/transaction/{id:[0-9]+}", r.handler.HandleGetTransactionByID).Methods("GET")
	r.router.HandleFunc("/transaction", r.handler.HandleCreateTransaction).Methods("POST")
	r.router.HandleFunc("/transaction/{id:[0-9]+}", r.handler.HandleUpdateTransaction).Methods("PUT")
	r.router.HandleFunc("/transaction/{id:[0-9]+}", r.handler.HandleDeleteTransaction).Methods("DELETE")
}
