package server

import (
	"context"
	"fmt"
	"net/http"

	"assets-liabilities/config"
	m "assets-liabilities/server/middlewares"
	"assets-liabilities/server/routes"
	"assets-liabilities/server/routes/finances"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server is object that manages routing and middleware
type Server struct {
	router *mux.Router
	server *http.Server
}

func (s *Server) bindRouter(ctx context.Context, r routes.Router) {
	// Each route can have multiple handlers associated with if the route supports multiple methods
	for route, methodHandlers := range r.List() {
		for method, handler := range methodHandlers {
			s.router.HandleFunc(route, handler).Methods(method, "OPTIONS")
		}
	}
}

// New performs initial setup of the server. Middlewares and routes are configued and new configuration and logging objects are created
func New(ctx context.Context) *Server {
	s := &Server{}

	s.router = mux.NewRouter()

	// Setup middleware
	s.router.Use(m.AddContext(ctx))
	s.router.Use(m.Logging(ctx))
	s.router.Use(handlers.CORS())

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{config.Config(ctx).AllowableOrigin})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// Configure routes
	// Not adding auth routes here
	// s.bindRouter(ctx, &auth.Router{})
	s.bindRouter(ctx, &finances.Router{})

	cfg := config.Config(ctx)

	s.server = &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(s.router),
		Addr:         fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		WriteTimeout: cfg.WriteTimeoutInSeconds,
		ReadTimeout:  cfg.ReadTimeoutInSeconds,
	}

	return s
}

// Start instructs the server to listen on the configured address and port
func (s *Server) Start() error {
	fmt.Print("Starting server\n")
	return s.server.ListenAndServe()
}
