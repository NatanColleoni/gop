package main

import (
	"github.com/NatanColleoni/gop/config"
	r "github.com/NatanColleoni/gop/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	r.Initialize()
}
