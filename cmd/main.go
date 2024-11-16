package main

import (
	"github.com/gin-gonic/gin"

	"github.com/serikdev/go-todo/internal/app/config"
	"github.com/serikdev/go-todo/internal/app/repository"
	"github.com/serikdev/go-todo/internal/app/routes"
	"github.com/serikdev/go-todo/internal/app/scheduler"
)

func main() {
	appConfig := config.InitCfg()

	repository.InitDB()
	r := gin.New()
	routes.InitRoutes(r)

	stopChan := make(chan struct{})
	go scheduler.OverdueUpdater(stopChan)

	if err := r.Run(appConfig.Address); err != nil {
		panic(err)
	}
	close(stopChan)

}
