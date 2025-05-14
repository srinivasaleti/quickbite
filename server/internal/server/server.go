package server

import (
	"fmt"
	"net/http"
	"quickbite/server/internal/product"
	"quickbite/server/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type IServer interface {
	Start(port string)
}

type Server struct {
	Logger logger.ILogger
	Port   string
}

func (s *Server) Start() {
	s.Logger.Info("starting server", "port", s.Port)
	r := s.handler()
	s.Logger.Info("started server", "port", s.Port)
	addr := ":" + s.Port
	err := http.ListenAndServe(addr, r)
	if err != nil {
		s.Logger.Error(err, "unable to start server", "port", s.Port)
		return
	}
}

func (c *Server) handler() *chi.Mux {
	r := chi.NewRouter()
	product := product.NewProductRouter()

	// all routes under /api
	r.Route("/api", func(api chi.Router) {
		product.AddRoutesToAppRouter(api)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	return r
}

func NewServer(port string) (*Server, error) {
	var s Server
	_logger, err := logger.NewLogger("info")
	if err != nil {
		fmt.Println("unable to create zap logger", "error", err)
		s.Logger = &logger.Logger{}
	}
	s.Logger = _logger
	s.Port = port
	s.Start()
	return &s, nil
}
