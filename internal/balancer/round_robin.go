package balancer

import (
	"fmt"
	"net/http"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

type RRLoadBalancer struct {
	port            int
	roundRobinCount int //server index to get next the request
	servers         []server.Servers
}

func (lb *RRLoadBalancer) Port() int {
	return lb.port
}
func (lb *RRLoadBalancer) GetNextAvailableServer() server.Servers {
	server := lb.servers[(lb.roundRobinCount)%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[(lb.roundRobinCount)%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *RRLoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Println("Request forwarded to:", targetServer.Address())
	targetServer.Serve(w, r)
}

func CreateLoadBalancers(port int, servers []server.Servers) *RRLoadBalancer {
	return &RRLoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}
