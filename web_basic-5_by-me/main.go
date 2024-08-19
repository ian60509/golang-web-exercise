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
	"user_server/config"
)



func initDB (cfg *config.Config) (db *gorm.DB) {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,cfg.Password,cfg.Network,cfg.Server,cfg.Port,cfg.Name) //設定資料庫連線資訊 (USERNAME
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //開啟資料庫連線 
	if err != nil {
		panic("use gorm connection MySQL failed, err:" + err.Error()) //如果錯誤引發panic
	}

	fmt.Println("connect MySQL success")

	if err := db.AutoMigrate(&model.User{}); err != nil { //會根據 User struct 來建立 table(或是比較現有的table 和 struct 的差異) => 如果有差異就會更新table
		panic("Migrate failure, err:" + err.Error())
	}
	fmt.Println("Migrate success")
	return db
}

func main() {
	config := config.NewConfig()
	db := initDB(config)

	userRepo := repository.NewUserRepository(db) //建立一個 userRepo struct
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


