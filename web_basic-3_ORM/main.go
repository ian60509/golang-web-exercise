package main

import (
	"gorm.io/driver/mysql" // gorm 提供的 mysql driver
	"gorm.io/gorm"
	"fmt"
)
  
const ( //database connection info 
	USERNAME = "demo"
	PASSWORD = "demo123"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "demo"
)
func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //開啟資料庫連線 
	if err != nil {
		panic("use gorm connection MySQL failed, err:" + err.Error()) //如果錯誤引發panic
	}

	fmt.Println("connect MySQL success")

	if err := db.AutoMigrate(&User{}); err != nil { //會根據 User struct 來建立 table(或是比較現有的table 和 struct 的差異) => 如果有差異就會更新table
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

	//-----------------更新資料-----------------
	user.Password = "test_pwd_update"
	if err := UpdateUser(db, user); err != nil {
		panic("Update user failure, err:" + err.Error())
	} else {
		user, _ = FindUser(db, user.ID)
		fmt.Println("Update user success: ", user)
	}

	//-----------------刪除資料-----------------
	if err := DeleteUser(db, user); err != nil {
		panic("Delete user failure, err:" + err.Error())
	} else {
		fmt.Println("Delete user success")
	}

}

type User struct { // 用來對應到資料庫中的table的物件
	ID int64 `json:"id" gorm:"primary_key;auto_increase"`
	Username string `json:"username"`
	Password string `json:"password"`
	Height int `json:"height"`
}

func CreateUser(db *gorm.DB, user *User) error {
	// 只需要 user 欄位和資料庫中的欄位對應即可
	return db.Create(user).Error //在這個database 新增一筆user 資料
	// 如果 Create 過程中有錯誤，會設置到 Error 屬性中
}

func FindUser(db *gorm.DB, id int64) (user *User, err error) { //回傳從資料庫中找到的這個物件
	user = new(User)
	user.ID = id
	err = db.First(user).Error //將這個物件直接丟給 gorm 幫忙找 => 尋找第一筆符合的資料
	return user, err
}

func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error //更新資料
}

func DeleteUser(db *gorm.DB, user *User) error {
	return db.Delete(user).Error //刪除資料
}