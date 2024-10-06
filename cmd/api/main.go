package main

import (
	"github.com/ciizo/go-gin-clean-arch/internal/user/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)

	router.Run("localhost:8080")
}
