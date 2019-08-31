package handler

import (
	"github.com/labstack/echo"
	"net/http"
)
type (
	IUser interface {
		List(e echo.Context) error
	}

	User struct {
		Id int
		Name string
	}
)

func NewUser() *User{
	return &User{}
	// return &IUser{}
}

func (*User) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "users")
	}
}
