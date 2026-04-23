// Package cache provides in-memory and disk-backed caching for the nasha gateway.
//
// Planned features:
//   - Thumbnail cache: generate and serve image/video thumbnails on demand.
//   - Directory listing cache: short-lived TTL to reduce driver round-trips.
//   - Cache invalidation on write / delete operations.
package cache

// Cache is the top-level cache manager.
// TODO: implement thumbnail generation (ffmpeg / imaging), LRU eviction, and TTL.
type Cache struct {
	// Dir is the directory on disk used to persist thumbnail files.
	Dir string
}

// New creates a Cache that stores persisted data under dir.
// TODO: initialise LRU eviction and background cleanup workers.
func New(dir string) *Cache {
	return &Cache{Dir: dir}
}
