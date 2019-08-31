package handler

import(
	"github.com/labstack/echo"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/chat_server2/model"
	"strconv"
	"fmt"
)

type (
	IMessage interface {
		List(c echo.Context)
	}
	Message struct {
		Id			int 			`json:"id"`
		Text		string 			`json:"text"`
		CreatedAt 	time.Time 		`json:"created_at"`
	}
)

var (
	conn, _ = dbr.Open("mysql", "root:@tcp(localhost:3306)/workout", nil)
	sess        = conn.NewSession(nil)
	// messages map[string]*Messages
)

func NewMessage() *Message{
	return &Message{}
}

func (*Message) List() echo.HandlerFunc {
	// return func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "messages")
	// }
	return nil
	
}

func (*Message) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return nil
		fmt.Println("エラー1")
	}
	message := model.GetMessage(c, id)
	if message == nil {
		return nil
		fmt.Println("エラー2")
	}
	return c.JSON(http.StatusOK, message)
}

func (*Message) Create(c echo.Context) error {
	return nil
}




