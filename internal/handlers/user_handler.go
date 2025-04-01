package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHanlder struct {
}

func NewUserHandler() *UserHanlder {
	return &UserHanlder{}
}

func (h *UserHanlder) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("/init", getUsers)
	}
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "init"})
}
