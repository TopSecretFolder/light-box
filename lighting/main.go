package main

import (
	"light-box/animation"
	"light-box/led"
	"net/http"

	"github.com/chromedp/cdproto/animation"
	"github.com/labstack/echo"
)

const VERSION = "v1.0.5"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Light Box Server: "+VERSION)
	})

	e.POST("/animation/push", func(ctx echo.Context) error {
		a := animation.Animation{}
		ctx.Bind(&a)
		animation.GlobalManager.Enqueue(a)
		return nil
	})

	go led.Loop()
	e.Logger.Fatal(e.Start(":80"))
}
