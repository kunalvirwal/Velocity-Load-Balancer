package server

import "net/http"

type Servers interface {

	// gives url of server
	Address() string

	// checks if the server is responding
	IsAlive() bool

	// serves the http request
	Serve(w http.ResponseWriter, r *http.Request)
}
