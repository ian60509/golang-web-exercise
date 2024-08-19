package main

import (
	"gorm.io/driver/mysql" // gorm 提供的 mysql driver
	"gorm.io/gorm"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"user_server/model"
)

const ( //database connection info 
	USERNAME = "demo"
	PASSWORD = "demo123"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "demo"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //開啟資料庫連線 
	if err != nil {
		panic("use gorm connection MySQL failed, err:" + err.Error()) //如果錯誤引發panic
	}

	fmt.Println("connect MySQL success")

	if err := DB.AutoMigrate(&model.User{}); err != nil { //會根據 User struct 來建立 table(或是比較現有的table 和 struct 的差異) => 如果有差異就會更新table
		panic("Migrate failure, err:" + err.Error())
	}
	fmt.Println("Migrate success")
}

func main() {
	server := gin.Default()
	userRoutes := server.Group("/user") //擁有相同prefix、middleware的routing group
	{
		userRoutes.GET("/:name", getUser)
		userRoutes.POST("", createUser)
		userRoutes.PUT("", updateUser)
		userRoutes.DELETE("/:name", deleteUser)
	}
	server.GET("/", healthCheck)
	server.Run(":8080")

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	// Fetch user data from database or any other source
	c.JSON(http.StatusOK, "get user name=" + name)
}

func createUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create user in database or any other source
	c.JSON(http.StatusOK, "create user" + user.Name)
}

func updateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update user in database or any other source
	c.JSON(http.StatusOK, "update user" + user.Name)
}

func deleteUser(c *gin.Context) {
	name := c.Param("name")
	// Delete user from database or any other source
	c.JSON(http.StatusOK, "delete user name=" + name)
}
