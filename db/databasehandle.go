package main

import (
    "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
)
// マイグレーションツールが見つかり使い方を習得し次第消す
func DatabaseInit() error {
	_, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/local_database")
	if err != nil{
		return err
	}
	return nil
}

// メモリアドレス不一致でパニックを起こした
func Migration() error{
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/local_database")
	if err != nil{
		return err
	}
	query, err := db.Query(`
		CREATE TABLE IF NOT EXISTS "messages" (
			"id" INTEGER, 
			"text" VARCHAR(255), 
			"created_at" DATETIME,
			PRIMARY KEY ("id")
		)`)
	defer query.Close()
	if err != nil{
		return err
	}
	return nil
}

func main() {
	DatabaseInit()
	Migration()
	defer fmt.Println("終了")
}