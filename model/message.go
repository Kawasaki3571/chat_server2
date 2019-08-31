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

func NewMessage2(id int, text string, createdAt time.Time) (*Message) {
	return &Message{Id: id,Text: text,CreatedAt: createdAt}
}

func GetMessage(c echo.Context, id int) (*Message) {
	db := GormConnect()
	defer db.Close()

	message := NewMessage(c)
	db.Find(&message, "id=?", id)
	return message
}
