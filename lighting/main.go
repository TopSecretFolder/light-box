package main

import (
	"fmt"
	"io"
	"light-box/animation"
	"light-box/led"
	"light-box/vue"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const VERSION = "v1.0.20"

func main() {
	e := echo.New()

	vueFS, err := vue.RawDistFS()
	if err != nil {
		log.Fatalf("error setting up static file handler: %w", err)
	}

	e.StaticFS("/", vueFS)

	e.GET("/about", func(c echo.Context) error {
		return c.String(http.StatusOK, "Light Box Server: "+VERSION)
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

	// The number of LEDs in your strip. We only use the first one.
	const NUM_LEDS = 18
	// The SPI device path. "/dev/spidev0.0" is standard.
	const spiPort = "/dev/spidev0.0" // An empty string uses the default SPI bus.

	strip, err := led.NewSK9822(spiPort, NUM_LEDS)
	if err != nil {
		log.Printf("Failed to initialize LED strip: %v\n", err)
	}
	// Ensure we clean up resources on exit.
	defer func() {
		if strip != nil {
			strip.Close()
		}
	}()

	go led.Loop(strip, NUM_LEDS)

	e.Logger.Fatal(e.Start(":80"))
}
