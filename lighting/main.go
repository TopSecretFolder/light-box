package main

import (
	"fmt"
	"light-box/animation"
	"light-box/led"
	"net/http"

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
		if err := ctx.Bind(&a); err != nil {
			return fmt.Errorf("error pushing animation: %w", err)
		}

		animation.GlobalManager.Enqueue(a)
		return nil
	})

	go led.Loop()
	e.Logger.Fatal(e.Start(":80"))
}
