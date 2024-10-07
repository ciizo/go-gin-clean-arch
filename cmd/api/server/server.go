package server

import (
	"github.com/ciizo/go-gin-clean-arch/internal/user/handlers"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	url    string
	router *gin.Engine
}

func New(url string) *ApiServer {
	g := gin.Default()

	return &ApiServer{url: url, router: g}
}

func (api *ApiServer) RegisterRoutes() {
	registerUserHandler(api.router)
}

func (api *ApiServer) Start() error {
	api.router.Run(api.url)
}

func registerUserHandler(router *gin.Engine) {
	router.GET("/users", handlers.GetUsers)
}
