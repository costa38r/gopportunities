package main

import (
	"github.com/costa38r/gopportunities/config"
	"github.com/costa38r/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()

}
