package server

import (
	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	mock.Mock
}

func (m *MockServer) Start() {
	m.Mock.Called()
}
