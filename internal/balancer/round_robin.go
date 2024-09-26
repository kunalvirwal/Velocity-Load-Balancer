package balancer

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

var mu sync.Mutex

type RRLoadBalancer struct {
	port            int
	RoundRobinCount int //server index to get next the request
	servers         []server.Servers
}

func (lb *RRLoadBalancer) Port() int {
	return lb.port
}

func (lb *RRLoadBalancer) GetAlgorythm() string {
	return "RoundRobin"
}

func (lb *RRLoadBalancer) GetNextAvailableServer() server.Servers {

	// for _, name := range lb.servers {
	// 	// fmt.Println(name.Address(), name.IsAlive())
	// }

	mu.Lock()
	server := lb.servers[(lb.RoundRobinCount)%len(lb.servers)]
	fmt.Println(*lb)
	for !server.IsAlive() {
		// fmt.Println(server.IsAlive())
		lb.RoundRobinCount++
		server = lb.servers[(lb.RoundRobinCount)%len(lb.servers)]
	} // TODO: Implement gracefull 503 responses if all servers down
	lb.RoundRobinCount++

	mu.Unlock()

	return server
}

func (lb *RRLoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	go targetServer.IncrementConnections()
	fmt.Println("Request forwarded to:", targetServer.Address())
	targetServer.Serve(w, r)
	go targetServer.DecrementConnections()

}
