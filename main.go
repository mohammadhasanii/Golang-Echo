package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Echo Framework")
	})

	app.Logger.Fatal(app.Start(":4000"))

}
