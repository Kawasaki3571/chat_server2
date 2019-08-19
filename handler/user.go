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

// func (h *News) List(c echo.Context) error {
// 	f, err := form.NewPaging(c)
// 	if err != nil {
// 		return err
// 	}
// 	newses, pagination, err := service.NewNews().GetList(f.Page, f.PerPage)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(200, map[string]interface{}{"news": response.NewNewses(newses), "pagination": response.NewPagination(pagination)})
// }