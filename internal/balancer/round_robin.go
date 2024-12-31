package balancer

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

type RRLoadBalancer struct {
	domain          string
	port            int
	RoundRobinCount int //server index to get next the request
	servers         []server.Servers
	mu              sync.Mutex
}

func (lb *RRLoadBalancer) Port() int {
	return lb.port
}

func (lb *RRLoadBalancer) GetAlgorythm() string {
	return "RoundRobin"
}

func (lb *RRLoadBalancer) GetNextAvailableServer() server.Servers {

	lb.mu.Lock()
	defer lb.mu.Unlock()
	server := lb.servers[(lb.RoundRobinCount)%len(lb.servers)]
	i := 0
	for !server.IsAlive() {
		i++
		if i == len(lb.servers) { // if all servers are down
			return nil
		}
		lb.RoundRobinCount++
		server = lb.servers[(lb.RoundRobinCount)%len(lb.servers)]
	}
	lb.RoundRobinCount++

	return server
}

func (lb *RRLoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	if targetServer == nil { // can redirect to a fallback server
		http.Error(w, "Service Unavailable: No healthy servers available", http.StatusServiceUnavailable)
		utils.LogNewError(fmt.Sprintf("Request Dropped %v: No healthy servers available", lb.domain))
		return
	}
	go targetServer.IncrementConnections()
	utils.Log(fmt.Sprintf("Request forwarded to: %v", targetServer.Address()))
	targetServer.Serve(w, r)
	go targetServer.DecrementConnections()

}
