package server

import (
	"fmt"
	"net/http"
	"quickbite/server/internal/logger"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server on port:", s.Port)
	})
	s.Logger.Info("started server", "port", s.Port)
	addr := ":" + s.Port
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		s.Logger.Error(err, "unable to start server", "port", s.Port)
		return
	}
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
