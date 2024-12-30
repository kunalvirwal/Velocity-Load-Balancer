package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/balancer"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/config"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/healthcheck"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

func main() {

	config.GetConfigs() // TODO : Implement global error handeling for invalid yaml
	const PORT = 8000
	var allServers []server.Servers
	LB := make(map[string]balancer.LoadBalancers)

	for _, service := range config.Cfgs.Services {
		fmt.Println(service.TargetURLs)
		BalancingAlgorythm := service.Algorythm
		var servers []server.Servers
		for _, url := range service.TargetURLs {
			server_instance := server.CreateServer(url)
			servers = append(servers, server_instance)
			allServers = append(allServers, server_instance) // TODO : can implement selective health checks here
		}
		// fmt.Println(servers)
		// fmt.Println(allServers)
		lb := balancer.CreateLoadBalancers(BalancingAlgorythm, PORT, servers) // "RoundRobin" // "LeastConnections" // Possible values present is ./internal/utils/utils.go
		if lb == nil {
			panic("Invalid balancing algorythm, nil load balancer recieved")
		}
		LB[service.Domain] = lb
	}

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		domain := req.Host
		lb, exist := LB[domain] // Get the load balancer for this service
		if !exist {
			http.Error(rw, "Service not found", http.StatusNotFound)
			fmt.Println("A request with unrecognised domain recieved, please update config.yml file or DNS ")
			return
		}
		lb.ServeProxy(rw, req)
	}

	go healthcheck.HealthCheck(allServers)

	// register a proxy handler to handle all requests
	http.HandleFunc("/", handleRedirect)

	for domain, lb := range LB {
		fmt.Println(lb.GetAlgorythm(), "loadbalancer serving requests at localhost:", (lb.Port()), "for the domain", (domain))
	}

	// TODO : run this (v) code only for distinct, it is possible that multiple lb run on same or different ports. Replace the hard coded 8000 port by lb.Port()
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		fmt.Println(err)
	}

}
