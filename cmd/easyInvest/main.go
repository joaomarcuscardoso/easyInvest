package main

import (
	"github.com/joaomarcuscardoso/easyInvest/internal/config"
	"github.com/joaomarcuscardoso/easyInvest/internal/router"
)

var (
	logger = *config.GetLogger("main")
)

func main() {
	logger = *config.GetLogger("main")

	// Initalize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
