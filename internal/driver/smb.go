// Package driver – SMB/CIFS driver stub.
// This file will grow into a full implementation backed by
// github.com/hirochachacha/go-smb2 for connecting to Samba / Windows shares.
package driver

import (
	"context"
	"io"

	"github.com/hirochachacha/go-smb2"
)

// SMBDriver implements StorageDriver for an SMB/CIFS share.
// TODO: implement all methods.
type SMBDriver struct {
	Host     string
	Share    string
	Username string
	Password string
	Domain   string

	// session is the authenticated SMB2 session; populated on first use.
	session *smb2.Session
}

func (d *SMBDriver) List(_ context.Context, _ string) ([]FileInfo, error) {
	panic("not implemented")
}

func (d *SMBDriver) Read(_ context.Context, _ string) (io.ReadCloser, error) {
	panic("not implemented")
}

func (d *SMBDriver) Write(_ context.Context, _ string, _ io.Reader) error {
	panic("not implemented")
}

func (d *SMBDriver) Delete(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *SMBDriver) MakeDir(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *SMBDriver) Rename(_ context.Context, _, _ string) error {
	panic("not implemented")
}

func (d *SMBDriver) Stat(_ context.Context, _ string) (FileInfo, error) {
	panic("not implemented")
}
