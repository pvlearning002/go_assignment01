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
}

type engine struct {
	app *gin.Engine
}

func NewEngine(cfg *config.Config) Engine {
	app := &engine{
		app: gin.Default(),
	}
	app.initRoutes(cfg)
	return app
}

func (e *engine) Start() error {
	cfg := config.GetConfig()
	return e.app.Run(cfg.AppPort) // Pass the configuration to the Run method
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.app.ServeHTTP(w, r)
}

func (e *engine) initRoutes(cfg *config.Config) {
	// Initialize gen pass service
	genPassService := service.NewGenPass()
	// Initialize gen pass handler
	genPassHandler := handler.NewGenPass(genPassService)
	// Register gen pass route
	e.app.GET(cfg.GEN_PASS_POST_PATH, genPassHandler.GeneratePassword)

	// Initialize health check service
	healthCheckService := service.NewHealthCheck()
	// Initialize health check handler
	healthCheckHandler := handler.NewHealthCheck(healthCheckService)
	// Register health check route
	e.app.GET(cfg., healthCheckHandler.GenerateHealthCheck)
}
