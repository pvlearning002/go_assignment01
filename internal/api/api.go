package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pvlearning002/go_assignment01/internal/config"
	"github.com/pvlearning002/go_assignment01/internal/handler"
	"github.com/pvlearning002/go_assignment01/internal/service"
)

type Engine interface {
	Start() error
	InitRoutes()
}

type engine struct {
	app *gin.Engine
}

func NewEngine() Engine {
	return &engine{
		app: gin.Default(),
	}
}

func (e *engine) Start() error {
	return e.app.Run(config.PORT_DEFAULT)
}

func (e *engine) InitRoutes() {
	healthCheckService := service.NewHealthCheck()
	healthCheckHandler := handler.NewHealthCheck(healthCheckService)
	e.app.GET(config.HEALTH_CHECK_GET_PATH, healthCheckHandler.GenerateHealthCheck)
}
