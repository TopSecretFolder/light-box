package main

import (
	"fmt"
	"light-box/animation"
	"light-box/led"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

const VERSION = "v1.0.11"

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
