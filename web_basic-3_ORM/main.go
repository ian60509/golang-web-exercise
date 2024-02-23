package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)
  
  const ( //database connection info 
	USERNAME = "root"
	PASSWORD = "22387681"
	NETWORK = "tcp"
	SERVER = "127.0.0.2"
	PORT = 3306
	DATABASE = "demo"
)
func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("use gorm connection MySQL failed, err:" + err.Error())
	}
	fmt.Println("connect MySQL success")

	if err := db.AutoMigrate(&User{}); err != nil {
		panic("Migrate failure, err:" + err.Error())
	}
	fmt.Println("Migrate success")

	//-----------------新增資料-----------------
	user := &User{
		Username: "test",
		Password: "test_pwd",
	}
	if err := CreateUser(db, user); err != nil {
		panic("Create user failure, err:" + err.Error())
	}
	fmt.Println("Create user success")

	//-----------------查詢資料-----------------
	if user, err = FindUser(db, user.ID); err == nil {
		fmt.Printf("Query user success, user:%#v\n", user)
	} else {
		panic("Query user failure, err:" + err.Error())
	}

}

type User struct {
	ID int64 `json:"id" gorm:"primary_key;auto_increase"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error //在這個database 新增一筆user 資料
}

func FindUser(db *gorm.DB, id int64) (user *User, err error) { //回傳從資料庫中找到的這個物件
	user = new(User)
	user.ID = id
	err = db.First(user).Error //尋找第一筆符合的資料
	return user, err
}