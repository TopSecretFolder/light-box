package main

import (
	"light-box/strip"
	"net/http"

	"github.com/labstack/echo"
)

const VERSION = "v1.0.5"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Light Box Server: "+VERSION)
	})

	go strip.Blink()

	e.Logger.Fatal(e.Start(":80"))
}
