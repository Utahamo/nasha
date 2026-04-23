// Package driver – SFTP driver stub.
// This file will grow into a full implementation backed by
// github.com/pkg/sftp over an SSH connection established via golang.org/x/crypto/ssh.
package driver

import (
	"context"
	"io"
)

// SFTPDriver implements StorageDriver for an SFTP server.
// TODO: implement all methods.
type SFTPDriver struct {
	Host       string
	Port       int
	Username   string
	Password   string
	// PrivateKey is the PEM-encoded private key used for key-based authentication.
	PrivateKey string
}

func (d *SFTPDriver) List(_ context.Context, _ string) ([]FileInfo, error) {
	panic("not implemented")
}

func (d *SFTPDriver) Read(_ context.Context, _ string) (io.ReadCloser, error) {
	panic("not implemented")
}

func (d *SFTPDriver) Write(_ context.Context, _ string, _ io.Reader) error {
	panic("not implemented")
}

func (d *SFTPDriver) Delete(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *SFTPDriver) MakeDir(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *SFTPDriver) Rename(_ context.Context, _, _ string) error {
	panic("not implemented")
}

func (d *SFTPDriver) Stat(_ context.Context, _ string) (FileInfo, error) {
	panic("not implemented")
}
