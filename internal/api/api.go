// Package api wires up the Fiber HTTP server and registers all route groups.
//
// Route layout (planned):
//
//	GET  /api/v1/health          – liveness probe
//	POST /api/v1/auth/login      – obtain JWT
//	POST /api/v1/auth/refresh    – refresh JWT
//	GET  /api/v1/fs/*path        – list directory / download file
//	PUT  /api/v1/fs/*path        – upload file
//	DELETE /api/v1/fs/*path      – delete file or directory
//	POST /api/v1/fs/mkdir        – create directory
//	POST /api/v1/fs/rename       – rename / move
//	GET  /api/v1/admin/users     – user management (RBAC protected)
//	*    /dav/*path              – WebDAV server endpoint (golang.org/x/net/webdav)
package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"golang.org/x/net/webdav"
)

// New creates and configures the Fiber application.
// TODO: wire in VFS, auth middleware, and individual route handlers.
func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "nasha",
	})

	// Health check – always available without authentication.
	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// TODO: register auth routes.
	// TODO: register filesystem routes (protected by JWT middleware).
	// TODO: register admin routes (protected by RBAC middleware).

	// WebDAV server endpoint – exposes the VFS as a WebDAV share.
	// TODO: wire up a real webdav.FileSystem backed by the VFS instead of MemFS.
	davHandler := &webdav.Handler{
		Prefix:     "/dav",
		FileSystem: webdav.NewMemFS(),
		LockSystem: webdav.NewMemLS(),
	}
	app.All("/dav/*", adaptor.HTTPHandler(davHandler))

	// Serve the React SPA for every non-API route.
	app.Static("/", "./web/dist")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/dist/index.html")
	})

	return app
}
