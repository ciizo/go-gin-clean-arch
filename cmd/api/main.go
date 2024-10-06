package main

import (
	"fmt"

	"github.com/ciizo/go-gin-clean-arch/config"
	"github.com/ciizo/go-gin-clean-arch/internal/user/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.New().LoadConfigurations()
	svUrl := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)

	router := gin.Default()
	router.GET("/users", handlers.GetUsers)

	router.Run(svUrl)
}
