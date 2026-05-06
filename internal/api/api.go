package api

import (
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/Utahamo/nasha/internal/auth"
	"github.com/Utahamo/nasha/internal/config"
	"github.com/Utahamo/nasha/internal/db"
	"github.com/Utahamo/nasha/internal/vfs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type API struct {
	v    *vfs.VFS
	db   *db.DB
	auth *auth.Auth
}

func New(v *vfs.VFS, database *db.DB, cfg *config.Config) *fiber.App {
	ttl, err := time.ParseDuration(cfg.Auth.TokenTTL)
	if err != nil {
		ttl = 24 * time.Hour
	}
	a := &API{
		v:  v,
		db: database,
		auth: auth.New(auth.Config{
			Secret:   cfg.Auth.JWTSecret,
			TokenTTL: ttl,
		}),
	}

	app := fiber.New(fiber.Config{AppName: "nasha"})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type",
	}))

	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Login — query DB, verify bcrypt
	app.Post("/api/v1/auth/login", func(c *fiber.Ctx) error {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}
		var user db.User
		if err := a.db.Where("username = ?", body.Username).First(&user).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
		}
		if user.Disabled {
			return c.Status(403).JSON(fiber.Map{"error": "account disabled"})
		}
		if !auth.CheckPassword(body.Password, user.PasswordHash) {
			return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
		}
		token, err := a.auth.SignToken(user.ID, user.Role)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to sign token"})
		}
		return c.JSON(fiber.Map{"token": token})
	})

	// File system routes (JWT required)
	fsGroup := app.Group("/api/v1/fs", a.auth.Middleware())

	// GET — list directory or read file
	fsGroup.Get("/*", a.handleGet)

	// POST — upload file(s) to path
	fsGroup.Post("/*", a.handlePost)

	// DELETE — delete file or directory
	fsGroup.Delete("/*", a.handleDelete)

	// PATCH — rename / move
	fsGroup.Patch("/*", a.handlePatch)

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

func (a *API) vpath(c *fiber.Ctx) string {
	return "/" + strings.TrimLeft(c.Params("*"), "/")
}

// GET — stat: file → send content, dir → list JSON
func (a *API) handleGet(c *fiber.Ctx) error {
	p := a.vpath(c)
	info, statErr := a.v.Stat(c.Context(), p)
	if statErr == nil && !info.IsDir {
		rc, err := a.v.Read(c.Context(), p)
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
	entries, err := a.v.List(c.Context(), p)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(entries)
}

// POST — multipart file upload or ?mkdir for directory creation
func (a *API) handlePost(c *fiber.Ctx) error {
	p := a.vpath(c)

	// Check if this is a mkdir request
	if c.Query("mkdir") == "1" {
		if err := a.v.MakeDir(c.Context(), p); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(201).JSON(fiber.Map{"status": "ok", "path": p})
	}

	// Otherwise multipart file upload
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid multipart form"})
	}
	var uploaded []string
	for _, headers := range form.File {
		for _, h := range headers {
			f, err := h.Open()
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "failed to read file: " + h.Filename})
			}
			dst := filepath.Join(p, h.Filename)
			if err := a.v.Write(c.Context(), dst, f); err != nil {
				f.Close()
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			f.Close()
			uploaded = append(uploaded, dst)
		}
	}
	return c.Status(201).JSON(fiber.Map{"status": "ok", "uploaded": uploaded})
}

// DELETE — remove file or directory
func (a *API) handleDelete(c *fiber.Ctx) error {
	p := a.vpath(c)
	if err := a.v.Delete(c.Context(), p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "ok"})
}

// PATCH — rename / move, body: {"src": "...", "dst": "..."}
func (a *API) handlePatch(c *fiber.Ctx) error {
	var body struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	}
	if err := c.BodyParser(&body); err != nil || body.Src == "" || body.Dst == "" {
		return c.Status(400).JSON(fiber.Map{"error": "src and dst required"})
	}
	if err := a.v.Rename(c.Context(), body.Src, body.Dst); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "ok"})
}
