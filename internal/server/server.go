package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

type Servers interface {

	// gives url of server
	Address() string

	// checks if the server is responding
	IsAlive() bool

	// serves the http request
	Serve(w http.ResponseWriter, r *http.Request)

	// Gets the no of active connections
	ActiveConnections() int

	// Increments the no of active connections
	IncrementConnections()

	// Decrements the no of active connections
	DecrementConnections()
}

func CreateServer(URL string) Servers { // TODO: Modify this function to accept type of server to create if needed
	serverURL, err := url.Parse(URL)
	utils.CheckNilErr(err, "Unable to parse url")

	return &Server{
		address:           URL,
		proxy:             httputil.NewSingleHostReverseProxy(serverURL),
		activeConnections: 0,
	}

}
