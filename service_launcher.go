package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/balancer"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/config"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/healthcheck"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/state"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
	"github.com/kunalvirwal/Velocity-Load-Balancer/pkg/api"
)

// creates all the load balancers for the services mentioned in the config file
func initLoadBalancers() {
	PORT := config.Cfgs.Listen_PORT
	for _, service := range config.Cfgs.Services {
		servers := createServers(&service)
		state.AllServers = append(state.AllServers, servers...)
		createAndRegisterLoadBalancer(&service, PORT, servers)
	}
	for domain, lb := range state.LB {
		utils.LogInfo(fmt.Sprintf("%v loadbalancer serving requests at localhost: %v for the domain %v", lb.GetAlgorythm(), lb.Port(), domain))
	}
}

// starts the request listener for the load balancers
func initListner() {

	PORT := config.Cfgs.Listen_PORT
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		domain := req.Host
		lb, exist := (state.LB)[domain] // Get the load balancer for this service
		if !exist {
			utils.Log("A request with unrecognised domain recieved, please update config.yml file or DNS ")
			http.Error(rw, "Service not found", http.StatusNotFound)
			return
		}
		lb.ServeProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil) // TODO : run this (v) code only for distinct, it is possible that multiple lb run on same or different ports. Replace the hard coded 8000 port by lb.Port()
	if err != nil {
		utils.LogError(err)
	} else {
		utils.LogInfo(fmt.Sprintf("API Listening on port %v", PORT))
	}

}

// starts the health check for the servers
func initHealthCheck() {
	healthcheck.HealthCheck(&state.AllServers)
}

func initAPI() {
	r := api.APIService()
	err := api.BuildHTTPServer(r, config.Cfgs.API_PORT).ListenAndServe()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
}

// creates all the servers for the service
func createServers(service *config.Service) []server.Servers {
	var servers []server.Servers
	for _, url := range (*service).TargetURLs {
		backend := server.CreateServer(url)
		servers = append(servers, backend)
	}
	return servers
}

// creates and registers the load balancer for a service
func createAndRegisterLoadBalancer(service *config.Service, port int, servers []server.Servers) {
	lb := balancer.CreateLoadBalancer(service.Domain, service.Algorythm, port, servers) // "RoundRobin" and "LeastConnections": Possible values present is ./internal/utils/utils.go
	if lb == nil {
		utils.LogNewError("Invalid balancing algorythm, nil load balancer recieved")
		os.Exit(1)
	}
	state.LB[service.Domain] = lb
}
