// routerはmainとは別に書く予定
package router
import(
	"github.com/labstack/echo"
	"github.com/chat_server2/handler"
)

func rooting(e *echo.Echo) {
	e.GET("/hello", handler.MainPage())
}