package model
import (
	// "github.com/chat_server2/model"
	"time"
	"github.com/labstack/echo"
)

type(
	Message struct {
		Id			int 			`json:"id"`
		Text		string 			`json:"text"`
		CreatedAt 	time.Time 		`json:"created_at"`
	}
)

func NewMessage(c echo.Context) (*Message) {
	return &Message{}
}

func GetMessage(c echo.Context, id int) (*Message) {
	db := gormConnect()
	defer db.Close()

	message := NewMessage(c)
	// return db.Find(&messages, "id=?", id)
	db.Find(&message, "id=?", id)
	return message
}
