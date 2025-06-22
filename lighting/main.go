package main

import (
	"fmt"
	"github.com/labstack/echo"
	"image"
	"image/color"
	"log"
	"net/http"

	"periph.io/x/conn/v3/display"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/apa102"
	"periph.io/x/devices/v3/screen1d"
	"periph.io/x/host/v3"
)

const VERSION = "v1.0.3"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Light Box Server: "+VERSION)
	})

	go lightLoop()

	e.Logger.Fatal(e.Start(":80"))
}

func lightLoop() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	d := getLEDs()
	img := image.NewNRGBA(d.Bounds())
	for x := range img.Rect.Max.X {
		img.SetNRGBA(x, 0, colorWheel(float64(x)/float64(img.Rect.Max.X)))
	}
	if err := d.Draw(d.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")
}

// getLEDs returns an *apa102.Dev, or fails back to *screen1d.Dev if no SPI port
// is found.
func getLEDs() display.Drawer {
	s, err := spireg.Open("")
	if err != nil {
		fmt.Printf("Failed to find a SPI port, printing at the console:\n")
		return screen1d.New(&screen1d.Opts{X: 100})
	}
	// Change the option values to see their effects.
	opts := apa102.DefaultOpts
	opts.NumPixels = 1
	d, err := apa102.New(s, &opts)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

// colorWheel returns a HSV color wheel.
func colorWheel(h float64) color.NRGBA {
	h *= 6
	switch {
	case h < 1.:
		return color.NRGBA{R: 255, G: byte(255 * h), A: 255}
	case h < 2.:
		return color.NRGBA{R: byte(255 * (2 - h)), G: 255, A: 255}
	case h < 3.:
		return color.NRGBA{G: 255, B: byte(255 * (h - 2)), A: 255}
	case h < 4.:
		return color.NRGBA{G: byte(255 * (4 - h)), B: 255, A: 255}
	case h < 5.:
		return color.NRGBA{R: byte(255 * (h - 4)), B: 255, A: 255}
	default:
		return color.NRGBA{R: 255, B: byte(255 * (6 - h)), A: 255}
	}
}
