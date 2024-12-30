package healthcheck

import (
	"fmt"
	"net"
	"time"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

func HealthCheck(servers []server.Servers) {
	ticker := time.NewTicker(10 * time.Second)

	// to start the health checks instantaneously
	for _, server := range servers {
		go runHealthChecks(server)
		// fmt.Println(server.IsAlive())

	}

	for range ticker.C {
		for _, server := range servers {
			go runHealthChecks(server)
			// fmt.Println(server.IsAlive())
		}
	}
}

func runHealthChecks(server server.Servers) {
	ConnectionTimeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", server.Address(), ConnectionTimeout)

	if err != nil {
		if server.IsAlive() {
			fmt.Println("HealthCheck:", server.Address(), "is offline")
			server.SetHealth(false)
			// fmt.Println(server.IsAlive())
		}
		return
	}

	defer conn.Close()

	if !server.IsAlive() {
		server.SetHealth(true)
		// fmt.Println(server.IsAlive())
		fmt.Println("HealthCheck:", server.Address(), "is now back online")
	}

}
