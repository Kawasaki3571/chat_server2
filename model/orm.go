package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
  DBMS     := "mysql"
  USER     := "root"
  PASS     := "root"
  PROTOCOL := "tcp(localhost:3306)"
  DBNAME   := "local_database"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}
