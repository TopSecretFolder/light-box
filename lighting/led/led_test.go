package led_test

import (
	"light-box/led"
	"testing"
)

func TestLoop(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		{"test1"},
	}
	for _, tt := range tests {
		const numLEDs = 1
		// The SPI device path. "/dev/spidev0.0" is standard.
		const spiPort = "/dev/spidev0.0" // An empty string uses the default SPI bus.

		strip, _ := led.NewSK9822(spiPort, numLEDs)

		t.Run(tt.name, func(t *testing.T) {
			led.Loop(strip)
		})
	}
}
