package server

import (
	"github.com/ciizo/go-gin-clean-arch/internal/user/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApiServer struct {
	address string
	router  *gin.Engine
	db      *gorm.DB
}

func New(url string, connectionString string) (*ApiServer, error) {
	g := gin.Default()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	return &ApiServer{url, g, db}, err
}

func (api *ApiServer) RegisterRoutes() {
	api.registerUserHandler()
}

func (api *ApiServer) Start() error {
	err := api.router.Run(api.address)

	return err
}

func (api *ApiServer) registerUserHandler() {
	h := handler.New(api.db)
	api.router.GET("/users", h.GetUsers)
}
