package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", HelloHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func HelloHandler(c echo.Context) error {
	fmt.Println("=>", c.Request().Host, "<=")
	return c.NoContent(http.StatusOK)
}

type User struct {
	Name string
}

func (u *User) Do() int {
	return 1
}
