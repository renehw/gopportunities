package main

import (
	"github.com/renehw/gopportunities.git/config"
	"github.com/renehw/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errf("config initialization failed: %v", err)
		return
	}

	// Inicializar o Router
	router.Initialize()
}
