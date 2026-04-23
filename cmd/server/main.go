// Command server is the nasha gateway binary.
// It loads configuration, wires up the VFS with configured storage drivers,
// starts the Fiber HTTP server, and handles graceful shutdown.
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Utahamo/nasha/internal/api"
)

func main() {
	app := api.New()

	addr := os.Getenv("NASHA_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	// Start the server in a goroutine so we can listen for OS signals.
	go func() {
		if err := app.Listen(addr); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	log.Printf("nasha listening on %s", addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("shutting down…")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
}
