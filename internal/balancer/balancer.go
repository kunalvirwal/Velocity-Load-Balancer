package balancer

import (
	"net/http"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

type LoadBalancers interface {
	// gets listeniong port of load balancer
	Port() int

	// gets the next server according to the algorythm used
	GetNextAvailableServer() server.Servers

	// forwards the request to the next server
	ServeProxy(w http.ResponseWriter, r *http.Request)
}
