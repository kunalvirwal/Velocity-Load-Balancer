package server

import (
	"net/http"
	"net/http/httputil"
	"sync"
)

type Baremetal_Server struct {
	address           string
	proxy             *httputil.ReverseProxy
	activeConnections int
	health            bool // health = true indicates that the serever is working properly
	mu                sync.RWMutex
}

// Gets address of Baremetal_Server
func (s *Baremetal_Server) Address() string {
	return s.address
}

// Checks if the Baremetal_Server is responding
func (s *Baremetal_Server) IsAlive() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.health
}

func (s *Baremetal_Server) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r) // TODO : add gracefull redirection on err
}

// Gets the no of active connections
func (s *Baremetal_Server) ActiveConnections() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.activeConnections
}

// Increments the no of active connections
func (s *Baremetal_Server) IncrementConnections() {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
}

// Decrements the no of active connections
func (s *Baremetal_Server) DecrementConnections() {
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Sets the health status of a Baremetal_Server
func (s *Baremetal_Server) SetHealth(health bool) {
	s.mu.Lock()
	s.health = health
	s.mu.Unlock()
}
