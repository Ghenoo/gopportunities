package main

import (
	"github.com/ghenoo/gopportunities/config"
	"github.com/ghenoo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Config
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initilization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
