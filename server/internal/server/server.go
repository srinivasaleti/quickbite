package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/srinivasaleti/quickbite/server/internal/config"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/domain/order"
	"github.com/srinivasaleti/quickbite/server/internal/domain/product"
	productsSeeder "github.com/srinivasaleti/quickbite/server/internal/domain/product/seeder"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
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
	dbConfig := s.Configuration.DBConfig()
	dbConfig.Logger = s.Logger
	database, err := database.NewDatabase(dbConfig)
	if err != nil {
		s.Logger.Error(err, "unable to connect to db")
		return
	}
	s.Logger.Info("successfully configured data store")
	defer database.Close()

	go s.seedData(database)
	r := s.handler(database)
	s.Logger.Info("started server", "port", s.Port)
	addr := ":" + s.Port
	err = http.ListenAndServe(addr, r)
	if err != nil {
		s.Logger.Error(err, "unable to start server", "port", s.Port)
		return
	}
}

func (s *Server) seedData(db database.DB) {
	productSeeder := productsSeeder.NewProductSeeder(s.Logger, db)
	productSeeder.SeedProducts()
}

func (s *Server) handler(db database.DB) *chi.Mux {
	r := chi.NewRouter()
	product := product.NewProductRouter(s.Logger, db)
	order := order.NewOrderRouter(s.Logger, db)

	if s.Configuration.Environment == config.LOCAL {
		r.Use(cors.AllowAll().Handler)
	}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// all routes under /api
	r.Route("/api", func(api chi.Router) {
		product.AddRoutesToAppRouter(api)
		order.AddRoutesToAppRouter(api)
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
