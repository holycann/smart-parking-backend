package payment_methods

import "github.com/gorilla/mux"

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
	r.router.HandleFunc("/payment_method", r.handler.HandleGetAllPaymentMethod).Methods("GET")
	r.router.HandleFunc("/payment_method/{id:[0-9]+}", r.handler.HandleGetPaymentMethodByID).Methods("GET")
	r.router.HandleFunc("/payment_method", r.handler.HandleCreatePaymentMethod).Methods("POST")
	r.router.HandleFunc("/payment_method/{id:[0-9]+}", r.handler.HandleUpdatePaymentMethod).Methods("PUT")
	r.router.HandleFunc("/payment_method/{id:[0-9]+}", r.handler.HandleDeletePaymentMethod).Methods("DELETE")
}
