package api

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/Utahamo/nasha/internal/auth"
	"github.com/Utahamo/nasha/internal/config"
	"github.com/Utahamo/nasha/internal/vfs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New(v *vfs.VFS, cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{AppName: "nasha"})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type",
	}))

	// Liveness probe
	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Login — hardcoded admin/admin123 for the demo
	app.Post("/api/v1/auth/login", func(c *fiber.Ctx) error {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}
		if body.Username != "admin" || body.Password != "admin123" {
			return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
		}
		token, err := auth.SignToken(1, "admin")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to sign token"})
		}
		return c.JSON(fiber.Map{"token": token})
	})

	// File system routes (JWT required)
	fsGroup := app.Group("/api/v1/fs", auth.Middleware())

	fsGroup.Get("/*", func(c *fiber.Ctx) error {
		vpath := "/" + strings.TrimLeft(c.Params("*"), "/")

		// Try to stat first; if it's a file, read and send its content
		info, statErr := v.Stat(c.Context(), vpath)
		if statErr == nil && !info.IsDir {
			rc, err := v.Read(c.Context(), vpath)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			defer rc.Close()
			data, err := io.ReadAll(rc)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			return c.Send(data)
		}

		// Otherwise list the directory
		entries, err := v.List(c.Context(), vpath)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(entries)
	})

	// SPA fallback
	staticDir := cfg.Server.StaticDir
	if staticDir == "" {
		staticDir = "./web/dist"
	}
	app.Static("/", staticDir)
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile(filepath.Join(staticDir, "index.html"))
	})

	return app
}
