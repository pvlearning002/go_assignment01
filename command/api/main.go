package main

import "github.com/pvlearning002/go_assignment01/internal/api"

func main() {
	engine := api.NewEngine()
	engine.InitRoutes()
	if err := engine.Start(); err != nil {
		panic(err)
	}
}
