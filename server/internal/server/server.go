package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/planner/server/internal/config"
	"github.com/srinivasaleti/planner/server/internal/database"
	"github.com/srinivasaleti/planner/server/internal/product"
	"github.com/srinivasaleti/planner/server/pkg/logger"
)

type IServer interface {
	Start()
}

type Server struct {
	Logger        logger.ILogger
	Port          string
	Configuration config.ServerConfiguration
}

func (s *Server) Start() {
	s.Logger.Info("starting server", "port", s.Port)
	// connecting to database
	s.Logger.Info("configuring data store")
	database, err := database.NewDatabase(s.Configuration.DBConfig())
	if err != nil {
		s.Logger.Error(err, "unable to connect to db")
		return
	}
	s.Logger.Info("successfully configured data store")
	defer database.Close()

	r := s.handler(s.Logger)
	s.Logger.Info("started server", "port", s.Port)
	addr := ":" + s.Port
	err = http.ListenAndServe(addr, r)
	if err != nil {
		s.Logger.Error(err, "unable to start server", "port", s.Port)
		return
	}
}

func (c *Server) handler(logger logger.ILogger) *chi.Mux {
	r := chi.NewRouter()
	product := product.NewProductRouter(logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// all routes under /api
	r.Route("/api", func(api chi.Router) {
		product.AddRoutesToAppRouter(api)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	return r
}

func NewServer(port string, config config.ServerConfiguration) (IServer, error) {
	var s Server
	_logger, err := logger.NewLogger("info")
	if err != nil {
		fmt.Println("unable to create zap logger", "error", err)
		s.Logger = &logger.Logger{}
	}
	s.Logger = _logger
	s.Port = port
	s.Configuration = config
	return &s, nil
}
