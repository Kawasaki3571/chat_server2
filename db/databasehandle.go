package main

import (
    "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
   "time"
)
// マイグレーションツールが見つかり使い方を習得し次第消す
func DatabaseInit() error {
	_, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/local_database")
	if err != nil{
		return err
	}
	return nil
}

// // メモリアドレス不一致でパニックを起こした
// func Migration() error{
// 	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
// 	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/local_database")
// 	if err != nil{
// 		return err
// 	}
// 	query, err := db.Query(`
// 		CREATE TABLE IF NOT EXISTS "messages" (
// 			"id" INTEGER, 
// 			"text" VARCHAR(255), 
// 			"created_at" DATETIME,
// 			PRIMARY KEY ("id")
// 		)`)
// 	defer query.Close()
// 	if err != nil{
// 		return err
// 	}
// 	return nil
// }
func Seed() error{
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/local_database")
	defer db.Close()
	if err != nil{
		return err
		fmt.Println("aaaaa!?")
	}
	_, err = db.Exec(`INSERT INTO messages VALUES ($1, $2, $3)`, 1, "test_message", time.Now())
	if err != nil {
		return err
		fmt.Println("aaaaa!?")
	}
	// defer query.Close()
	return nil
}

func main() {
	DatabaseInit()
	// Migration()
	Seed()
	defer fmt.Println("終了")
}

