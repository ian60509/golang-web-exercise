package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql" // gorm 提供的 mysql driver
	"gorm.io/gorm"

	"user_server/model"
	"user_server/module/delivery"
	"user_server/module/service"
	"user_server/module/repository"
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

func initDB() {
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
	initDB()

	userRepo := repository.NewUserRepository(DB) //建立一個 userRepo struct
	userService := service.NewUserService(userRepo) //建立一個 userService struct
	userHandler := delivery.NewUserHandler(userService) //建立一個 userHandler struct

	server := gin.Default()
	userRoutes := server.Group("/user") //擁有相同prefix、middleware的routing group
	{
		userRoutes.GET("/:name", userHandler.GetUser)
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.PUT("", userHandler.UpdateUser)
		userRoutes.DELETE("/:name", userHandler.DeleteUser)
	}
	server.GET("/", userHandler.HealthCheck)
	server.Run(":8080")

}


