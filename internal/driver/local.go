// Package driver – local filesystem driver stub.
// This file will grow into a full implementation that mounts a local directory
// as a nasha storage endpoint.
package driver

import (
	"context"
	"io"
)

// LocalDriver implements StorageDriver for a local filesystem path.
// TODO: implement all methods.
type LocalDriver struct {
	// Root is the absolute path to the directory exposed by this driver.
	Root string
}

func (d *LocalDriver) List(_ context.Context, _ string) ([]FileInfo, error) {
	panic("not implemented")
}

func (d *LocalDriver) Read(_ context.Context, _ string) (io.ReadCloser, error) {
	panic("not implemented")
}

func (d *LocalDriver) Write(_ context.Context, _ string, _ io.Reader) error {
	panic("not implemented")
}

func (d *LocalDriver) Delete(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *LocalDriver) MakeDir(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *LocalDriver) Rename(_ context.Context, _, _ string) error {
	panic("not implemented")
}

func (d *LocalDriver) Stat(_ context.Context, _ string) (FileInfo, error) {
	panic("not implemented")
}
