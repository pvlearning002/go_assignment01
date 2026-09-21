package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvlearning002/go_assignment01/internal/config"
	"github.com/pvlearning002/go_assignment01/internal/handler"
	service "github.com/pvlearning002/go_assignment01/internal/service"
)

type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	InitRoutes()
}

type engine struct {
	app *gin.Engine
}

func NewEngine() Engine {
	app := &engine{
		app: gin.Default(),
	}
	return app
}

func (e *engine) Start() error {
	return e.app.Run(config.PORT_DEFAULT)
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.app.ServeHTTP(w, r)
}

func (e *engine) InitRoutes() {
	// Initialize gen pass service
	genPassService := service.NewGenPass()
	// Initialize gen pass handler
	genPassHandler := handler.NewGenPass(genPassService)
	// Register gen pass route
	e.app.GET(config.GEN_PASS_POST_PATH, genPassHandler.GeneratePassword)

	// Initialize health check service
	healthCheckService := service.NewHealthCheck()
	// Initialize health check handler
	healthCheckHandler := handler.NewHealthCheck(healthCheckService)
	// Register health check route
	e.app.GET(config.HEALTH_CHECK_GET_PATH, healthCheckHandler.GenerateHealthCheck)
}
