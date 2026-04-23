// Package driver defines the StorageDriver interface and FileInfo type.
// Every storage backend (local, WebDAV, SMB, S3, SFTP, …) must implement
// StorageDriver so that higher-level layers (VFS, API) remain protocol-agnostic.
package driver

import (
	"context"
	"io"
	"time"
)

// FileInfo describes a file or directory returned by a storage driver.
type FileInfo struct {
	Name    string
	Size    int64
	IsDir   bool
	ModTime time.Time
	Path    string
}

// StorageDriver is the core abstraction for all storage backends.
// All methods accept a context so that callers can cancel long-running
// operations (e.g. large file transfers).
type StorageDriver interface {
	// List returns the contents of the directory at the given path.
	List(ctx context.Context, path string) ([]FileInfo, error)

	// Read opens the file at the given path for reading.
	// Callers are responsible for closing the returned ReadCloser.
	Read(ctx context.Context, path string) (io.ReadCloser, error)

	// Write streams data from r into the file at the given path,
	// creating intermediate directories as needed.
	Write(ctx context.Context, path string, r io.Reader) error

	// Delete removes the file or empty directory at the given path.
	Delete(ctx context.Context, path string) error

	// MakeDir creates the directory at the given path (and any missing parents).
	MakeDir(ctx context.Context, path string) error

	// Rename moves or renames the file/directory from src to dst.
	Rename(ctx context.Context, src, dst string) error

	// Stat returns metadata for the file or directory at the given path.
	Stat(ctx context.Context, path string) (FileInfo, error)
}
