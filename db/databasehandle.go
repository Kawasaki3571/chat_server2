package main

import (
    "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
)
// マイグレーションツールが見つかり使い方を習得し次第消す
func Migration() error{
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
	if err != nil{
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "messages" ("id" INTEGER, "text" VARCHAR(255), "created_at" DATETIME)`,)
	if err != nil{
		return err
	}
	return nil
}

func main() {
	Migration()
	defer fmt.Println("終了")
}