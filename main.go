package main

import (
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/config"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/healthcheck"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

func main() {

	config.GetConfigs() // TODO : Implement global error handeling for invalid yaml
	var allServers []server.Servers
	var LB LoadBalancers = make(LoadBalancers)

	initLoadBalancers(config.Cfgs.Listen_PORT, &LB, &allServers)
	go healthcheck.HealthCheck(allServers)
	initListner(config.Cfgs.Listen_PORT, &LB)

}
