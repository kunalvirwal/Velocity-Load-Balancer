package balancer

import (
	"net/http"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

type LoadBalancers interface {

	// Gets the port on which the load balancer is running
	Port() int

	// gets the next server according to the algorythm used
	GetNextAvailableServer() server.Servers

	// forwards the request to the next server
	ServeProxy(w http.ResponseWriter, r *http.Request)

	// gets the algorythm being used by that load balancer
	GetAlgorythm() string
}

// create a new load baklancer instance
func CreateLoadBalancers(algorythm string, port int, servers []server.Servers) LoadBalancers {

	if algorythm == utils.RRenum {
		return &RRLoadBalancer{
			port:            port,
			RoundRobinCount: 0,
			servers:         servers,
		}
	} else if algorythm == utils.LCenum {
		return &LCLoadBalancer{
			port:    port,
			servers: servers,
		}
	}

	return nil
}
