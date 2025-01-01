package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/config"
	c "github.com/kunalvirwal/Velocity-Load-Balancer/pkg/controllers"
)

// create service with servers
// add server to service
// remove server from service
// delete service
// get service
// all services
// update algorythm for service

// create api request & response types
// implement mutual uniqueness in yaml file
// controllers defined here would be used with gRPC also

func APIService() *gin.Engine {

	r := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{fmt.Sprintf("http://localhost:%v", config.Cfgs.API_PORT)},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
		MaxAge:          12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	service := r.Group("/service")
	{
		service.GET("", c.GetAllServices)
		// service.GET("/:service", c.getService)
		// service.POST("/create", c.createNewService)
		// service.PUT("/:service/addserver", c.addServerToService)
		// service.DELETE("/:service", c.deleteService)
	}
	// server := r.Group("/server")
	// {
	// 	server.DELETE("/:service", c.removeServerFromService)
	// 	server.DELETE("", c.deleteServerFromAllServices)
	// }

	r.NoRoute(c.Handle404)

	return r

}

func BuildHTTPServer(r *gin.Engine, p int) *http.Server {

	port := strconv.Itoa(p)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    5 * time.Second,  // Maximum duration for reading the entire request
		WriteTimeout:   10 * time.Second, // Maximum duration before timing out writes of the response
		IdleTimeout:    10 * time.Second, // Maximum amount of time to wait for the next request when keep-alives are enabled
		MaxHeaderBytes: 1 << 20,
	}

	return s
}
