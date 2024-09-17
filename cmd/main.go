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
		server.CreateServer("https://www.duckduckgo.com"),
		server.CreateServer("https://www.facebook.com"),
		server.CreateServer("https://www.bing.com"),
	}

	lb := balancer.CreateLoadBalancers(8000, servers)

	fmt.Println(lb.Port())

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	// register a proxy handler to handle all requests
	http.HandleFunc("/", handleRedirect)

	fmt.Println("serving requests at localhost: ", (lb.Port()))
	http.ListenAndServe(":"+strconv.Itoa(lb.Port()), nil)

}
