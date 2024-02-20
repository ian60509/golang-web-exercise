package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)
var  UserData map[string]string
type IndexData struct { //定義好 template 內所要的參數
	Title string
	Content string
}

func init () {
	UserData = map[string]string{
		"test": "test",
	}
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
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
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