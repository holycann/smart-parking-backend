package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/notifications"
	payment_methods "github.com/holycann/smart-parking-backend/internal/payment_method"
	"github.com/holycann/smart-parking-backend/internal/reservations"
	"github.com/holycann/smart-parking-backend/internal/spots"
	"github.com/holycann/smart-parking-backend/internal/transactions"
	"github.com/holycann/smart-parking-backend/internal/users"
	vehicles "github.com/holycann/smart-parking-backend/internal/vehicle"
	"github.com/holycann/smart-parking-backend/internal/zones"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userModule(s, subrouter)
	notificationModule(s, subrouter)
	paymentMethodModule(s, subrouter)
	reservationModule(s, subrouter)
	spotModule(s, subrouter)
	transactionModule(s, subrouter)
	vehicleModule(s, subrouter)
	zoneModule(s, subrouter)

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	handlerWithCORS := corsMiddleware(router)

	log.Print("Listening On Port ", s.addr)

	return http.ListenAndServe(s.addr, handlerWithCORS)
}

func userModule(s *Server, router *mux.Router) {
	userRepository := users.NewRepository(s.db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)
	users.NewRoutes(router, userHandler)
}

func notificationModule(s *Server, router *mux.Router) {
	notificationRepository := notifications.NewRepository(s.db)
	notificationService := notifications.NewService(notificationRepository)
	notificationHandler := notifications.NewHandler(notificationService)
	notifications.NewRoutes(router, notificationHandler)
}

func paymentMethodModule(s *Server, router *mux.Router) {
	paymentMethodRepository := payment_methods.NewRepository(s.db)
	paymentMethodService := payment_methods.NewService(paymentMethodRepository)
	paymentMethodHandler := payment_methods.NewHandler(paymentMethodService)
	payment_methods.NewRoutes(router, paymentMethodHandler)
}

func reservationModule(s *Server, router *mux.Router) {
	reservationRepository := reservations.NewRepository(s.db)
	reservationService := reservations.NewService(reservationRepository)
	reservationHandler := reservations.NewHandler(reservationService)
	reservations.NewRoutes(router, reservationHandler)
}

func spotModule(s *Server, router *mux.Router) {
	spotRepository := spots.NewRepository(s.db)
	spotService := spots.NewService(spotRepository)
	spotHandler := spots.NewHandler(spotService)
	spots.NewRoutes(router, spotHandler)
}

func transactionModule(s *Server, router *mux.Router) {
	transactionRepository := transactions.NewRepository(s.db)
	transactionService := transactions.NewService(transactionRepository)
	transactionHandler := transactions.NewHandler(transactionService)
	transactions.NewRoutes(router, transactionHandler)
}

func vehicleModule(s *Server, router *mux.Router) {
	vehicleRepository := vehicles.NewRepository(s.db)
	vehicleService := vehicles.NewService(vehicleRepository)
	vehicleHandler := vehicles.NewHandler(vehicleService)
	vehicles.NewRoutes(router, vehicleHandler)
}

func zoneModule(s *Server, router *mux.Router) {
	zoneRepository := zones.NewRepository(s.db)
	zoneService := zones.NewService(zoneRepository)
	zoneHandler := zones.NewHandler(zoneService)
	zones.NewRoutes(router, zoneHandler)
}
