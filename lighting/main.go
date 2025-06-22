package main

import (
	"github.com/labstack/echo"
	"net/http"
)

const VERSION = "v1.0.1"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Light Box Server: "+VERSION)
	})
	e.Logger.Fatal(e.Start(":80"))
}
