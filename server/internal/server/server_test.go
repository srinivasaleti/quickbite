package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/config"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestServer_Start(t *testing.T) {
	server, _ := NewServer("18081", config.ServerConfiguration{})

	t.Run("should correctly handle routes", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:8081/health", nil)
		rr := httptest.NewRecorder()
		handler := server.(*Server).handler(&database.PostgresDB{})
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 400 when path not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:8081/invalid", nil)
		rr := httptest.NewRecorder()
		handler := server.(*Server).handler(&database.PostgresDB{})
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

}
