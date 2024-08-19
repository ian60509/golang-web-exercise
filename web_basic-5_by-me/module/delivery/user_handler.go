package delivery

import (
	"fmt"
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
	user, GetUserErr := h.service.GetUserByName(name)
	if GetUserErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": GetUserErr.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	
	// 從 request 的 JSON 中解析成 user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Bind JSON error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CretaeErr := h.service.Create(&user)
	if CretaeErr != nil {
		fmt.Println("Create user error: ", CretaeErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": CretaeErr.Error()})
		return
	}
	
	c.JSON(http.StatusOK, "create user: " + user.Name)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	UpdateErr := h.service.UpdateUser(&user)
	if UpdateErr != nil {
		fmt.Println("Update user error: ", UpdateErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": UpdateErr.Error()})
		return
	}
	c.JSON(http.StatusOK, "update user" + user.Name)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	name := c.Param("name")
	// Delete user from database or any other source

	DeleteErr := h.service.DeleteUser(name)
	if DeleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": DeleteErr.Error()})
	}
	c.JSON(http.StatusOK, "delete user name=" + name)
}