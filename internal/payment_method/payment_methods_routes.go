package payment_methods

import (
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type PaymentMethodRoutes struct {
	router  *mux.Router
	handler *PaymentMethodHandler
}

func NewRoutes(router *mux.Router, handler *PaymentMethodHandler) *PaymentMethodRoutes {
	return &PaymentMethodRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *PaymentMethodRoutes) RegisterRoutes() {
	router := r.router.PathPrefix("/payment_method").Subrouter()

	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", r.handler.HandleGetAllPaymentMethod).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleGetPaymentMethodByID).Methods("GET")
	router.HandleFunc("", r.handler.HandleCreatePaymentMethod).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleUpdatePaymentMethod).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", r.handler.HandleDeletePaymentMethod).Methods("DELETE")
}
