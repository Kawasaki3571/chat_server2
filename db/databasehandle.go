package main

import (
    "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
   "time"
   "github.com/chat_server2/model"
)
// マイグレーションツールが見つかり使い方を習得し次第消す
func DatabaseInit() error {
	_, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/local_database")
	if err != nil{
		return err
	}
	return nil
}

func Seed() error {
	db := model.GormConnect()
	defer db.Close()

	message := model.NewMessage2(1, "test_pyo", time.Now())
	db.Create(&message)
	return nil
}

func main() {
	DatabaseInit()
	Seed()
	defer fmt.Println("終了")
}

