package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model //這樣的定義使得 User 結構體自動擁有 ID、CreatedAt、UpdatedAt 和 DeletedAt 字段，並且這些字段會在數據庫操作中自動處理，而你只需要關注業務邏輯相關的字段（如 Name、Email 和 Password）。
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}