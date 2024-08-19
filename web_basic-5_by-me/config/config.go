package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Username string //database username
	Password string //database password
	Network string //ex: tcp
	Server string //database server IP address
	Port int //database server IP address
	Name string //database database name
}

func NewConfig() (cfg *Config) {
	viper.SetConfigName("config") // 指定讀取的檔名
	viper.AddConfigPath("config") // 指定讀取的路徑
	viper.SetConfigType("yaml") // 指定讀取的檔案格式
	
	err := viper.ReadInConfig() // 指定讀取的檔案
	if err != nil {
		fmt.Printf("Fatal error config file: %v \n", err)
	}
	cfg = &Config{
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Network: viper.GetString("database.network"),
		Server: viper.GetString("database.server"),
		Port: viper.GetInt("database.port"),
		Name: viper.GetString("database.name"),
	}
	return cfg
}