package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ciizo/go-gin-clean-arch/internal/user/models"
)

func (h *handler) GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

var users = []models.User{
	{ID: "1", Title: "Mr.", Name: "John", Age: 20},
	{ID: "2", Title: "Mr.", Name: "Mac", Age: 30},
	{ID: "3", Title: "Miss", Name: "Sara", Age: 40},
}
