package httpserver

import (
	"fmt"
	"io"
	"light-box/animation"
	"light-box/vue"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewEchoServer(version string) *echo.Echo {

	e := echo.New()

	vueFS, err := vue.RawDistFS()
	if err != nil {
		log.Fatalf("error setting up static file handler: %s", err.Error())
	}

	e.StaticFS("/", vueFS)
	e.GET("/about", func(c echo.Context) error {
		return c.String(http.StatusOK, "Light Box Server: "+version)
	})
	e.POST("/animation/push/json", func(ctx echo.Context) error {
		ani := animation.Animation{}
		if err := ctx.Bind(&ani); err != nil {
			return fmt.Errorf("error pushing animation: %w", err)
		}

		animation.GlobalManager.Enqueue(ani)
		return ctx.JSON(200, ani)
	})
	e.POST("/animation/push/lua", func(ctx echo.Context) error {
		b, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return fmt.Errorf("error pushing animation: %w", err)
		}

		ani, err := animation.ResolveLua(string(b))
		if err != nil {
			return err
		}

		animation.GlobalManager.Enqueue(ani)

		return ctx.JSON(200, ani)
	})
	e.POST("/script/push/lua", func(ctx echo.Context) error {
		// TODO: register script as current frame generator
		_, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return fmt.Errorf("error pushing animation: %w", err)
		}

		// push a lua script for frame gen

		return ctx.JSON(200, nil)
	})

	return e
}
