package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Utahamo/nasha/internal/api"
	"github.com/Utahamo/nasha/internal/config"
	"github.com/Utahamo/nasha/internal/driver"
	"github.com/Utahamo/nasha/internal/vfs"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	v := vfs.New()
	for _, mc := range cfg.Mounts {
		if mc.Type != "local" {
			log.Printf("skipping non-local mount %q (type=%s) — only local supported in demo", mc.Name, mc.Type)
			continue
		}
		v.Mount(&vfs.MountPoint{
			Name:   mc.Name,
			Path:   mc.Path,
			Driver: &driver.LocalDriver{Root: mc.Config["root"]},
		})
		log.Printf("mounted %q at %s → %s", mc.Name, mc.Path, mc.Config["root"])
	}

	app := api.New(v, cfg)

	addr := os.Getenv("NASHA_ADDR")
	if addr == "" {
		addr = cfg.Server.Addr
	}
	if addr == "" {
		addr = ":8080"
	}

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
