package strip

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

// SK9822 represents our LED strip
type SK9822 struct {
	conn    spi.Conn
	port    spi.PortCloser
	numLEDs int
	buffer  []byte
}

// NewSK9822 initializes the LED strip and its buffer.
func NewSK9822(portName string, numLEDs int) (*SK9822, error) {
	// Initialize the host drivers for hardware access.
	if _, err := host.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize periph host: %w", err)
	}

	// Open the specified SPI port. An empty string "" opens the default port.
	port, err := spireg.Open(portName)
	if err != nil {
		return nil, fmt.Errorf("failed to open SPI port '%s': %w", portName, err)
	}

	// Connect to the SPI device with a specific configuration.
	// SK9822s can handle high speeds. 8MHz is a safe starting point.
	// We use the physic.MegaHertz constant for clarity and correctness.
	// Mode0 is standard for this type of LED. Data is sent in 8-bit words.
	conn, err := port.Connect(8*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		port.Close()
		return nil, fmt.Errorf("failed to connect to SPI device: %w", err)
	}

	// The data buffer needs space for:
	// 1. A 4-byte start frame (all zeros).
	// 2. A 4-byte frame for each LED.
	// 3. A 4-byte end frame (all ones) to ensure data is clocked out.
	bufferSize := (numLEDs + 2) * 4

	s := &SK9822{
		conn:    conn,
		port:    port,
		numLEDs: numLEDs,
		buffer:  make([]byte, bufferSize),
	}

	// Initialize the buffer structure.
	// The start frame is at the beginning of the buffer. Since a new slice is
	// zero-initialized, we don't need to explicitly set it to 0x00000000.

	// The end frame is at the end of the buffer.
	endFrameOffset := (numLEDs + 1) * 4
	for i := 0; i < 4; i++ {
		s.buffer[endFrameOffset+i] = 0xFF
	}

	return s, nil
}

// SetPixel updates the color for a single LED in the local buffer.
// It takes the LED's index, RGB color values (0-255), and a brightness value (0-31).
func (s *SK9822) SetPixel(index int, r, g, b byte, brightness byte) {
	if index < 0 || index >= s.numLEDs {
		return // Ignore out-of-bounds requests.
	}
	// The SK9822 brightness field is 5 bits, so the max value is 31.
	if brightness > 31 {
		brightness = 31
	}

	// The actual LED data starts after the 4-byte start frame.
	offset := (index + 1) * 4

	// Each LED frame consists of 4 bytes:
	// Byte 1: 3-bit header (111) followed by 5-bit global brightness.
	// Byte 2: Blue channel.
	// Byte 3: Green channel.
	// Byte 4: Red channel.
	s.buffer[offset] = 0b11100000 | brightness
	s.buffer[offset+1] = b
	s.buffer[offset+2] = g
	s.buffer[offset+3] = r
}

// Render transmits the entire buffer to the LED strip via SPI.
func (s *SK9822) Render() error {
	// The Tx function performs the SPI transaction.
	// We send our buffer and don't need to receive any data back, so the read buffer is nil.
	return s.conn.Tx(s.buffer, nil)
}

// Clear turns off all LEDs in the buffer.
func (s *SK9822) Clear() {
	for i := 0; i < s.numLEDs; i++ {
		s.SetPixel(i, 0, 0, 0, 0)
	}
}

// Close gracefully shuts down the SPI port and turns off LEDs.
func (s *SK9822) Close() {
	log.Println("Closing SPI port and turning off LEDs.")
	s.Clear()
	s.Render()
	s.port.Close()
}

func Blink() {
	// The number of LEDs in your strip. We only use the first one.
	const numLEDs = 1
	// The SPI device path. "/dev/spidev0.0" is standard.
	const spiPort = "/dev/spidev0.0" // An empty string uses the default SPI bus.

	strip, err := NewSK9822(spiPort, numLEDs)
	if err != nil {
		log.Fatalf("Failed to initialize LED strip: %v", err)
	}
	// Ensure we clean up resources on exit.
	defer strip.Close()

	// Set up a channel to listen for OS signals (like Ctrl+C) for a graceful shutdown.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		strip.Close()
		os.Exit(0)
	}()

	log.Println("Blinking pixel 0. Press Ctrl+C to exit.")

	// Infinite loop to blink the pixel.
	for {
		// --- Turn the pixel ON (Blue) ---
		log.Println("Pixel ON")
		// Set the first pixel (index 0) to blue with a moderate brightness (10 out of 31).
		strip.SetPixel(0, 0, 0, 255, 10)
		// Send the data to the strip.
		if err := strip.Render(); err != nil {
			log.Printf("Failed to render ON state: %v", err)
		}
		time.Sleep(time.Second)

		// --- Turn the pixel OFF (Black) ---
		log.Println("Pixel OFF")
		// Set the first pixel to black (all zeros).
		strip.SetPixel(0, 0, 0, 0, 0)
		// Send the data to the strip.
		if err := strip.Render(); err != nil {
			log.Printf("Failed to render OFF state: %v", err)
		}
		time.Sleep(time.Second)
	}
}
