package main

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/chat_server2/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/joho/godotenv"
	"github.com/chat_server2/twitter"
	"log"

)

var (
	conn, _ = dbr.Open("mysql", "root:@tcp(localhost:3306)/workout", nil)
	sess        = conn.NewSession(nil)
)
func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}


func main() {
	loadEnv()
	e := echo.New()
	go twitter.GetTweet()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/hello", handler.MainPage())
	router.Rooting(e)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Info("shutting down the server")
	}
}
