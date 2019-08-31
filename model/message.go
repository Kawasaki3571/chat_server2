package model
import (
	"github.com/chat_server2/model"
)

type(
	Message struct {
		Id			int 			`json:"id"`
		Text		string 			`json:"text"`
		CreatedAt 	time.Time 		`json:"created_at"`
	}
)

func NewMessage(c echo.Context) (*Message) {
	return &NewMessage{}
}

func GetMessage(c echo.Context, id) (*Message) {
	db := model.gormConnect()
	defer db.Close()

	messages := NewMessage()
	return db.Find(&messages, "id=?", id) 
}