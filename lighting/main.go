package main

import (
	"light-box/httpserver"
	"light-box/led"
	"light-box/natsserver"
	"log"
)

const VERSION = "v1.0.20"

func main() {
	// NATS SETUP
	natsserver.SetupGlobalNatsManager()

	// LED CONTROLLER PROCESS
	const NUM_LEDS = 18
	const spiPort = "/dev/spidev0.0"
	strip, err := led.NewSK9822(spiPort, NUM_LEDS)
	if err != nil {
		log.Printf("Failed to initialize LED strip: %v\n", err)
	}
	defer func() {
		if strip != nil {
			strip.Close()
		}
	}()
	go led.Loop(strip, NUM_LEDS)

	// HTTP PROCESS
	e := httpserver.NewEchoServer(VERSION)
	e.Logger.Fatal(e.Start(":80"))
}
