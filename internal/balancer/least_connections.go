package balancer

import (
	"fmt"
	"net/http"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

type LCLoadBalancer struct {
	port    int
	servers []server.Servers
}

func (lb *LCLoadBalancer) Port() int {
	return lb.port
}

func (lb *LCLoadBalancer) GetAlgorythm() string {
	return "LeastConnections"
}

func (lb *LCLoadBalancer) GetNextAvailableServer() server.Servers {
	i := 0 // iterator to find first active server
	server := lb.servers[i]
	for !server.IsAlive() {
		i++
		if i == len(lb.servers) {
			// time.Sleep(1 * time.Second) // timeout to retry to check if any of the server comes alive
			// i = 0
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
	}
	targetServer.IncrementConnections()
	fmt.Println("Request forwarded to:", targetServer.Address())
	targetServer.Serve(w, r)
	targetServer.DecrementConnections()
}
