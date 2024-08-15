package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main () {
	server := gin.Default() //建立 gin instance
	server.LoadHTMLGlob("template/html/*") //載入 templates 資料夾內的所有檔案
	server.Static("/assets", "./template/assets") //設定靜態檔案路由，讓HTML可以直接透過 "/assets" 引用 "./template/assets" 資料夾內的檔案
	server.GET("/login", LoginPage) //將/login這個path 綁訂到LoginPage這個function
	server.POST("/login", LoginAuth) //將/login這個path(如果面臨 http POST request) 綁訂到LoginAuth這個function
	server.Run(":8080") //啟動 server (預設 port 為 8080
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil) //顯示登入頁面的HTML，因為前面有透過 LoadHTMLGlob 載入 templates 資料夾內的所有檔案，所以這邊可以直接寫檔名
}

