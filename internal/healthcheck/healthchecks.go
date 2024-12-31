package healthcheck

import (
	"fmt"
	"net"
	"time"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/server"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

func HealthCheck(servers []server.Servers) {
	ticker := time.NewTicker(10 * time.Second)

	// to start the health checks instantaneously
	for _, server := range servers {
		go runHealthChecks(server)
	}

	for range ticker.C {
		for _, server := range servers {
			go runHealthChecks(server)
		}
	}
}

func runHealthChecks(server server.Servers) {
	ConnectionTimeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", server.Address(), ConnectionTimeout)

	if err != nil {
		if server.IsAlive() {
			utils.LogCustom(utils.Red, "Healthcheck", fmt.Sprintf(" %v is offline", server.Address()))
			server.SetHealth(false)
		}
		return
	}

	defer conn.Close()

	if !server.IsAlive() {
		server.SetHealth(true)
		utils.LogCustom(utils.Green, "Healthcheck", fmt.Sprintf(" %v is now back online", server.Address()))
	}

}
