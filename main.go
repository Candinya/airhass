package main

import (
	"airhass/global"
	"airhass/inits"
	"log"
)

func readConfing() {

}

func main() {
	inits.Hello()

	// Initialize configuration
	if err := inits.Config(); err != nil {
		log.Fatalf("Failed to initialize config with error: %v", err)
	}

	// Initialize logger
	if err := inits.Logger(); err != nil {
		log.Fatalf("Failed to initialize logger with error: %v", err)
	}

	// Switch logger
	global.Logger.Debug("Logger initialized, switch to here.")

	// Initialize serial port
	if err := inits.Serial(); err != nil {
		global.Logger.Fatalf("Failed to initialize serial port with error: %v", err)
	}

	// Prepare to clean serial port when finish program
	defer global.SerialPort.Close()

	// Start waiting for events
	if err := inits.Watch(); err != nil {
		global.Logger.Fatalf("Failed to watch with error: %v", err)
	}

}
