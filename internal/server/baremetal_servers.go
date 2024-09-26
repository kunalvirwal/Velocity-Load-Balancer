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
	health            bool // health = true indicates that the serever is working properly
	mu                sync.RWMutex
}

// server functions
func (s *Server) Address() string {
	return s.address
}

// Checks if the server is responding
func (s *Server) IsAlive() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.health
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r) // TODO : add gracefull redirection on err
}

// Gets the no of active connections
func (s *Server) ActiveConnections() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
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

// Sets the health status of a server
func (s *Server) SetHealth(health bool) {
	s.mu.Lock()
	s.health = health
	s.mu.Unlock()
}
