// routerはmainとは別に書く予定
package router
import(
	"github.com/labstack/echo"
	"github.com/chat_server2/handler"
)

func Rooting(e *echo.Echo) {
	e.GET("/hello", handler.MainPage())

	user := handler.NewUser()
	e.GET("/users", user.List())

	message := handler.NewMessage()
	e.GET("/messages", message.List())
} 