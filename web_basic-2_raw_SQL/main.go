package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "22387681"
	NETWORK = "tcp"
	SERVER = "127.0.0.2"
	PORT = 3306
	DATABASE = "demo"
)

type User struct {
	ID string
	Username string
	Password string
}


func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE) //格式化連線字串
	db, err := sql.Open("mysql", conn) //開啟此連線之資料庫
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return
	} else {
		fmt.Println("開啟 MySQL 連線成功")
	}

    if err := db.Ping(); err != nil { //檢查資料庫連線
        fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return
	}
    defer db.Close()

	CreateTable(db)
	InsertUser(db, "test", "test")
	QuryUser(db, "test")
}

func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS users (
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64)
		); `

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("建立資料表發生錯誤，原因為：", err)
		return err
	}
	fmt.Println("建立Table users 成功")
	return nil
}

func InsertUser(db *sql.DB, username, pwd string) error {
	_,err := db.Exec("insert INTO users(username,password) values(?,?)",username, pwd) //插入一組 username, password 的資料

	if err != nil {
		fmt.Println("新增使用者發生錯誤，原因為：", err)
		return err
	}
	fmt.Println("新增使用者成功")
	return nil
}

func QuryUser(db *sql.DB, username string)  {
	user := new(User)
	row := db.QueryRow("select * from users where username=?", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil { //Scan( )顺序将查询结果中的列值依次读取到 user.ID、user.Username 和 user.Password 变量中，如果數量、類型不匹配會回傳錯誤
		fmt.Println("查詢使用者發生錯誤，原因為：", err)
		return
	}
	fmt.Println("查詢使用者成功，使用者資料為：", *user)
}