package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/save", save)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.POST("/save", save)
func save(c echo.Context) error {
	var name string
	c.Bind(&name)
	return c.String(http.StatusOK, "name:"+name)
}
