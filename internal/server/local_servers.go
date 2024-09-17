package server

import (
	"net/http"
	"net/http/httputil"
	"sync"
)

type Server struct {
	address           string
	proxy             *httputil.ReverseProxy
	activeConnections int
	mu                sync.Mutex
}

// server functions
func (s *Server) Address() string {
	return s.address
}

func (s *Server) IsAlive() bool {
	// TODO: write logic
	return true
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

// Gets the no of active connections
func (s *Server) ActiveConnections() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.activeConnections
}

// Increments the no of active connections
func (s *Server) IncrementConnections() {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
}

// Decrements the no of active connections
func (s *Server) DecrementConnections() {
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}
