// Package db manages the metadata database used by nasha.
// It wraps GORM with a SQLite backend for local metadata storage
// (mount configuration, user accounts, permissions, share links, etc.).
//
// Planned features:
//   - Auto-migration of all model tables on startup.
//   - Helper constructors for User, Mount, ShareLink models.
//   - Support for swapping SQLite → PostgreSQL via DSN prefix.
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB wraps a GORM database connection.
type DB struct {
	*gorm.DB
}

// Open opens (or creates) the SQLite database at the given DSN path
// and returns a ready-to-use DB handle.
// TODO: run AutoMigrate for all models once they are defined.
func Open(dsn string) (*DB, error) {
	g, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{g}, nil
}
