package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ciizo/go-gin-clean-arch/cmd/api/server"
	"github.com/ciizo/go-gin-clean-arch/config"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New().LoadConfigurations()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Cannot create logger", zap.Error(err))
	}

	srvUrl := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)
	srv := server.New(srvUrl)
	srv.RegisterRoutes()

	err = srv.Start()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("Cannot start api server", zap.Error(err))
	}
}
