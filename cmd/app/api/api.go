package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/holycann/smart-parking-backend/internal/users"
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
