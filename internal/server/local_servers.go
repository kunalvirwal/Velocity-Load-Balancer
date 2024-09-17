package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
)

type Server struct {
	address string
	proxy   *httputil.ReverseProxy
}

// server functions
func (s *Server) Address() string {
	return s.address
}

func (s *Server) IsAlive() bool {
	// TODO: write logic
	return true
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

func CreateServer(URL string) *Server {
	serverURL, err := url.Parse(URL)
	utils.CheckNilErr(err, "Unable to parse url")

	return &Server{
		address: URL,
		proxy:   httputil.NewSingleHostReverseProxy(serverURL),
	}

}
