package main

import (
	"backend_go/config"
	"backend_go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLooger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
