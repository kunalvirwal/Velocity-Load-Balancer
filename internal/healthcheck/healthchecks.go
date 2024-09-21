package healthcheck

import (
	"fmt"
	"net"
	"time"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
)

func HealthCheck(servers []server.Servers) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		for _, server := range servers {
			runHealthChecks(server)

		}
	}
}

func runHealthChecks(server server.Servers) {
	ConnectionTimeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", server.Address(), ConnectionTimeout)

	if err != nil {
		// fmt.Println(err)
		fmt.Println("HealthCheck:", server.Address(), "is offline")
		server.SetHealth(false)
		return
	}

	defer conn.Close()
	if !server.IsAlive() {
		server.SetHealth(true)
		fmt.Println("HealthCheck:", server.Address(), "is now back online")
	}

}
