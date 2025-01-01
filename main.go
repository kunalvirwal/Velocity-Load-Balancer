package main

import (
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/config"
)

func initServices() {
	go initHealthCheck()
	go initListner()
	go initAPI()
}

func main() {

	config.GetConfigs()
	initLoadBalancers()
	initServices()

	<-make(chan struct{})
}
