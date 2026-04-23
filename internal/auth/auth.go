// Package auth provides JWT-based authentication and RBAC authorisation.
//
// Planned features:
//   - Issue signed JWT access tokens on successful login.
//   - Refresh token flow with short-lived access tokens.
//   - Fiber middleware that validates the Authorization: Bearer header.
//   - Role-Based Access Control: admin, editor, viewer roles.
//   - Per-mount permission overrides stored in the database (via GORM).
package auth

import (
	"github.com/gofiber/fiber/v2"
)

// Claims are the custom JWT payload fields.
// TODO: extend with roles and per-mount permissions.
type Claims struct {
	UserID uint   `json:"uid"`
	Role   string `json:"role"`
}

// Middleware returns a Fiber handler that validates the JWT in the
// Authorization header and injects the parsed Claims into the request context.
// TODO: implement token parsing and validation using golang-jwt/jwt.
func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: extract and verify JWT, set claims in locals.
		return c.Next()
	}
}
