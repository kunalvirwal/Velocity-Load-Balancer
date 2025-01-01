package state

import (
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/balancer"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

// LoadBalancers is a map of domain to the load balancer for that domain
type LoadBalancers map[string]balancer.LoadBalancers

var AllServers []server.Servers
var LB LoadBalancers = make(LoadBalancers)
