package service

import(
	"time"
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

func Create(c echo.Context) (*Message) {
	m := NewMessage(c)
	
}