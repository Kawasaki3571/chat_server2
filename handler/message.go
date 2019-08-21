package handler

import(
	"github.com/labstack/echo"
	"net/http"
	"github.com/chat_server2/service"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
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

var (
	conn, _ = dbr.Open("mysql", "root:@tcp(localhost:3306)/workout", nil)
	sess        = conn.NewSession(nil)
)


func NewMessage() *Message{
	return &Message{}
}

func (*Message) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "messages")
	}
}

func (*Message) Create(c echo.Context) error {
	message := service.NewMessage(c)
	sess.InsertInto("messageinfo").Columns("id", "text", "created_at").Values(message.Id, message.Text, time.Now()).Exec()
	return c.NoContent(http.StatusOK)
}
// func InsertAuthor(c echo.Context) error {
//     author := new(Author)
//     if err := c.Bind(author); err != nil {
//         return err
//     }
//     sess.InsertInto(authortable).Columns("id", "name").Values(author.Id, author.Name).Exec()
//     return c.NoContent(http.StatusOK)
// }