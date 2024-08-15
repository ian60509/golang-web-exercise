package main

import (
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)
var  UserData map[string]string
type IndexData struct { //定義好 template 內所要的參數
	Title string
	Content string
}

// package 初始化時就會執行
func init () {
	UserData = map[string]string{
		"test": "test",
	}

	fmt.Println("Init UserData: ")
	for k, v := range UserData {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
	fmt.Println("--------------\n\n\n")
}

func CheckUserIsExist(name string) bool {
	if _, ok := UserData[name]; ok {
		return true
	}
	return false
}

func CheckPwd(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	}
	return errors.New("password is not correct")
}

func Auth(username, InputPwd string) error {
	if userExist := CheckUserIsExist(username); !userExist {
		return errors.New("user not exist")
	} else { //user exist
		return CheckPwd(UserData[username], InputPwd)
	}
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)


	// -------- 確認 username 欄位有輸入 --------
	if in, exist := c.GetPostForm("username"); exist  { //使用gin提供的表單功能
		username = in 
	} else { //沒有填寫 username
		// 這個 gin.H 並非header，而是一個用於傳遞資料的map[string]interface{}配合gin的HTML模板引擎
		c.HTML(http.StatusBadRequest, "login.html", gin.H{ //用於配合HTML的"模板引擎"，將error資料傳入HTML，讓 HTML 中透過 .error 捕捉到訊息 
			"error": errors.New("必須輸入使用者名稱"),
		}) //將錯誤訊息放在 header 回傳，這個header欄位是一個 ket-value pair => key: error, value: "必須輸入使用者名稱"
		return
	}

	// -------- 確認 password 欄位有輸入 --------
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}

	// -------- 驗證使用者名稱與密碼 --------
	if err := Auth(username, password); err != nil {
		// 設定登入成功的訊息，這個訊息會在HTML中的 .success 捕捉到
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"success": "登入成功啦啦啦", //回傳成功訊息
		})
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": err, //將這個驗證的錯誤訊息回傳
		})
	}
}