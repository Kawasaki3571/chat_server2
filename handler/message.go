package handler

import(
	"github.com/labstack/echo"
	"net/http"
)

type (
	IMessage interface {
		List(c echo.Context)
	}
	Message struct{
		// Id		int
		// Text	string
	}
)

func NewMessage() *Message{
	return &Message{}
}

func (*Message) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "messages")
	}
}