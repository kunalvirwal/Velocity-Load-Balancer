package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/balancer"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/config"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

// LoadBalancers is a map of domain to the load balancer for that domain
type LoadBalancers map[string]balancer.LoadBalancers

// CreateAllLoadBalancers creates all the load balancers for the services mentioned in the config file
func initLoadBalancers(PORT int, LB *LoadBalancers, allServers *[]server.Servers) {
	for _, service := range config.Cfgs.Services {
		utils.LogInfo(fmt.Sprintf("%v : %v", service.Domain, service.TargetURLs))
		var servers []server.Servers

		for _, url := range service.TargetURLs {

			backend := server.CreateServer(url)
			servers = append(servers, backend)
			*allServers = append(*allServers, backend) // TODO : can implement selective health checks here

		}
		// "RoundRobin" and "LeastConnections": Possible values present is ./internal/utils/utils.go
		lb := balancer.CreateLoadBalancer(service.Domain, service.Algorythm, PORT, servers)
		if lb == nil {
			utils.LogNewError("Invalid balancing algorythm, nil load balancer recieved")
			os.Exit(1)
		}
		(*LB)[service.Domain] = lb
	}

	for domain, lb := range *LB {
		utils.LogInfo(fmt.Sprintf("%v loadbalancer serving requests at localhost: %v for the domain %v", lb.GetAlgorythm(), lb.Port(), domain))
	}
}

func initListner(PORT int, LB *LoadBalancers) {
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		domain := req.Host
		lb, exist := (*LB)[domain] // Get the load balancer for this service
		if !exist {
			http.Error(rw, "Service not found", http.StatusNotFound)
			utils.Log("A request with unrecognised domain recieved, please update config.yml file or DNS ")
			return
		}
		lb.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)
	// TODO : run this (v) code only for distinct, it is possible that multiple lb run on same or different ports. Replace the hard coded 8000 port by lb.Port()
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		utils.LogError(err)
	}
}
