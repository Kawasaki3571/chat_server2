package main

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/chat_server2/handler"
	"github.com/chat_server2/router"
	// "github.com/chat_server2/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

var (
	conn, _ = dbr.Open("mysql", "root:@tcp(localhost:3306)/workout", nil)
	sess        = conn.NewSession(nil)
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/hello", handler.MainPage())
	router.Rooting(e)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Info("shutting down the server")
	}
}
