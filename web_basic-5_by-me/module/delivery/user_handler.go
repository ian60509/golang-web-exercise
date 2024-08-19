package delivery

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"user_server/model"
	"user_server/module/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler { //帶有各個很多 method 的 struct
	return &UserHandler{
		service: service,
	}
}

// ----------------------- method -----------------------------------

func (h *UserHandler)HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	name := c.Param("name")
	// Fetch user data from database or any other source
	c.JSON(http.StatusOK, "get user name=" + name)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create user in database or any other source
	c.JSON(http.StatusOK, "create user" + user.Name)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update user in database or any other source
	c.JSON(http.StatusOK, "update user" + user.Name)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	name := c.Param("name")
	// Delete user from database or any other source
	c.JSON(http.StatusOK, "delete user name=" + name)
}