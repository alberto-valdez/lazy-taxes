package handlers

import (
	"net/http"

	"github.com/alberto-valdez/lazy-taxes/internal/models"
	"github.com/alberto-valdez/lazy-taxes/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHanlder struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHanlder {
	return &UserHanlder{
		userService: userService,
	}
}

func (h *UserHanlder) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("/create", h.handleCreateUser)
		users.GET("/init", h.handleGetUser)
	}
}

func (h *UserHanlder) handleGetUser(c *gin.Context) {
	data, err := h.userService.GetUser()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *UserHanlder) handleCreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
