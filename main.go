package main

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/chat_server2/handler"
	"github.com/chat_server2/router"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/hello", handler.MainPage())
	router.rooting(e)


	// go func() {
	if err := e.Start(":1323"); err != nil {
		e.Logger.Info("shutting down the server")
	}
	// }()
}
