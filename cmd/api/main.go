package main

import (
	"fmt"

	"github.com/ciizo/go-gin-clean-arch/cmd/api/server"
	"github.com/ciizo/go-gin-clean-arch/config"
)

func main() {
	cfg := config.New().LoadConfigurations()
	srvUrl := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)

	srv := server.New(srvUrl)
	srv.RegisterRoutes()

	srv.Start()
}
