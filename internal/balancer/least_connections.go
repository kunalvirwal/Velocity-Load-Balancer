package balancer

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

type LCLoadBalancer struct {
	domain  string
	port    int
	servers []server.Servers
	mu      sync.Mutex
}

func (lb *LCLoadBalancer) Port() int {
	return lb.port
}

func (lb *LCLoadBalancer) GetAlgorythm() string {
	return "LeastConnections"
}

func (lb *LCLoadBalancer) GetNextAvailableServer() server.Servers {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	i := 0 // iterator to find first active server
	server := lb.servers[i]
	for !server.IsAlive() {
		i++
		if i == len(lb.servers) { // if all servers are down
			i = 0
			return nil
		}
		server = lb.servers[i]
	}
	for _, backend := range lb.servers {
		if backend.IsAlive() && (backend.ActiveConnections() < server.ActiveConnections()) {
			server = backend
		}
	}

	return server
}

func (lb *LCLoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	if targetServer == nil { // can redirect to a fallback server
		http.Error(w, "Service Unavailable: No healthy servers available", http.StatusServiceUnavailable)
		utils.LogNewError(fmt.Sprintf("Request Dropped %v: No healthy servers available", lb.domain))
		return
	}
	go targetServer.IncrementConnections()
	targetServer.Serve(w, r)
	go targetServer.DecrementConnections()
}
