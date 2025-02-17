package notifications

import "github.com/gorilla/mux"

type NotificationRoutes struct {
	router  *mux.Router
	handler *NotificationHandler
}

func NewRoutes(router *mux.Router, handler *NotificationHandler) *NotificationRoutes {
	return &NotificationRoutes{
		router:  router,
		handler: handler,
	}
}

func (r *NotificationRoutes) RegisterRoutes() {
	r.router.HandleFunc("/notification", r.handler.HandleGetAllNotifications).Methods("GET")
	r.router.HandleFunc("/notification/{id:[0-9]+}", r.handler.HandleGetAllNotifications).Methods("GET")
	r.router.HandleFunc("/notification", r.handler.HandleCreateNotification).Methods("POST")
	r.router.HandleFunc("/notification/{id:[0-9]+}", r.handler.HandleUpdateNotification).Methods("PUT")
	r.router.HandleFunc("/notification/{id:[0-9]+}", r.handler.HandleDeleteNotification).Methods("DELETE")
}
