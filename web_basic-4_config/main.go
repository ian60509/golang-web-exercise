package main

import (
	"github.com/spf13/viper"
	"fmt"
)

func main() {
	viper.SetConfigName("app") // 設定檔名稱 (不需要帶副檔名)
	viper.SetConfigType("yaml") // 設定檔案類型
	viper.AddConfigPath(".") // 設定檔案路徑
	viper.SetDefault("application.port", 8080) // 設定預設值。 某些參數如果沒有填寫可能會造成程式運行出現錯誤，因此我們可以透過 SetDefault 將特定參數設定預設值
	

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	fmt.Println("讀取成功")

	// 注意欄位名稱要按照yaml 階層結構來設定 : application.port
	fmt.Println(viper.Get("application.port"))
	fmt.Println(viper.Get("application.mode"))
}