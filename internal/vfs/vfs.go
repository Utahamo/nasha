// Package vfs implements the Virtual File System layer.
// It manages mount points and routes incoming paths to the correct
// StorageDriver, hiding the underlying protocol from the REST API layer.
//
// Planned responsibilities:
//   - Maintain a registry of MountPoint → StorageDriver mappings.
//   - Resolve a virtual path to the correct driver + relative path.
//   - Provide unified List/Read/Write/… operations that delegate to drivers.
//   - Handle cross-driver operations (e.g. server-side copy across two mounts).
package vfs

import "github.com/Utahamo/nasha/internal/driver"

// MountPoint associates a virtual path prefix with a storage driver.
type MountPoint struct {
	// Name is the human-readable label shown in the UI.
	Name string
	// Path is the virtual mount path (e.g. "/nas", "/cloud").
	Path string
	// Driver handles all file operations for this mount.
	Driver driver.StorageDriver
}

// VFS is the root of the virtual file system.
// TODO: implement routing, path resolution, and cross-driver operations.
type VFS struct {
	mounts []*MountPoint
}

// New creates an empty VFS with no mounts.
func New() *VFS {
	return &VFS{}
}

// Mount registers a new storage backend under the given path prefix.
// TODO: validate that path does not conflict with existing mounts.
func (v *VFS) Mount(mp *MountPoint) {
	v.mounts = append(v.mounts, mp)
}
