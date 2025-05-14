package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	mock.Mock
}

func (m *MockServer) Start() {
	m.Mock.Called()
}

func (m *MockServer) Handler() *chi.Mux {
	return chi.NewMux()
}
