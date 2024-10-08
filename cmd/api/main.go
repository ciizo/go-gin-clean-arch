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
		log.Fatal("Failed to create logger", zap.Error(err))
	}

	fmt.Printf("connectionstring: %s", cfg.Db.ConnectionString)

	srvUrl := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)
	srv, err := server.New(srvUrl, cfg.Db.ConnectionString)
	if err != nil {
		logger.Fatal("Failed to create api server", zap.Error(err))
	}

	srv.RegisterRoutes()

	err = srv.Start()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("Failed to start api server", zap.Error(err))
	}
}
