package main

import (
	"github.com/gabs1337/gopportunities/config"
	"github.com/gabs1337/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
