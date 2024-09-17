package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/balancer"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

func main() {
	servers := []server.Servers{
		server.CreateServer("http://localhost:8001"),
		server.CreateServer("http://localhost:8002"),
		server.CreateServer("http://localhost:8003"),
	}

	BalancingAlgorythm := "LeastConnections" // "RoundRobin"  // Possible values present is ./internal/utils/utils.go

	lb := balancer.CreateLoadBalancers(BalancingAlgorythm, 8000, servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	if lb == nil {
		fmt.Println("Invalid Balancing Algorythm chosen")
	} else {
		// register a proxy handler to handle all requests
		http.HandleFunc("/", handleRedirect)

		fmt.Println(BalancingAlgorythm, "loadbalancer serving requests at localhost:", (lb.Port()))
		http.ListenAndServe(":"+strconv.Itoa(lb.Port()), nil)
	}

}
