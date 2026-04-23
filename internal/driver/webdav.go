// Package driver – WebDAV driver stub.
// This file will grow into a full implementation that proxies operations to a
// remote WebDAV server (e.g. Nextcloud, ownCloud).
// Dependency: golang.org/x/net/webdav (client mode via stdlib http.Client).
package driver

import (
	"context"
	"io"
)

// WebDAVDriver implements StorageDriver for a WebDAV remote.
// TODO: implement all methods.
type WebDAVDriver struct {
	// Endpoint is the full URL of the WebDAV share (e.g. https://dav.example.com/remote.php/dav/files/user/).
	Endpoint string
	Username string
	Password string
}

func (d *WebDAVDriver) List(_ context.Context, _ string) ([]FileInfo, error) {
	panic("not implemented")
}

func (d *WebDAVDriver) Read(_ context.Context, _ string) (io.ReadCloser, error) {
	panic("not implemented")
}

func (d *WebDAVDriver) Write(_ context.Context, _ string, _ io.Reader) error {
	panic("not implemented")
}

func (d *WebDAVDriver) Delete(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *WebDAVDriver) MakeDir(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *WebDAVDriver) Rename(_ context.Context, _, _ string) error {
	panic("not implemented")
}

func (d *WebDAVDriver) Stat(_ context.Context, _ string) (FileInfo, error) {
	panic("not implemented")
}
