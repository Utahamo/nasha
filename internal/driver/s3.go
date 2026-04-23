// Package driver – S3-compatible object storage driver stub.
// This file will grow into a full implementation backed by
// github.com/aws/aws-sdk-go-v2/service/s3, supporting AWS S3, MinIO, Backblaze B2, etc.
package driver

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Driver implements StorageDriver for an S3-compatible object store.
// TODO: implement all methods.
type S3Driver struct {
	Endpoint        string
	Bucket          string
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	// PathStyle forces path-style addressing (required by MinIO and some other S3-compatible stores).
	PathStyle bool

	// client is the configured AWS S3 client; populated on first use.
	client *s3.Client
}

func (d *S3Driver) List(_ context.Context, _ string) ([]FileInfo, error) {
	panic("not implemented")
}

func (d *S3Driver) Read(_ context.Context, _ string) (io.ReadCloser, error) {
	panic("not implemented")
}

func (d *S3Driver) Write(_ context.Context, _ string, _ io.Reader) error {
	panic("not implemented")
}

func (d *S3Driver) Delete(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *S3Driver) MakeDir(_ context.Context, _ string) error {
	panic("not implemented")
}

func (d *S3Driver) Rename(_ context.Context, _, _ string) error {
	panic("not implemented")
}

func (d *S3Driver) Stat(_ context.Context, _ string) (FileInfo, error) {
	panic("not implemented")
}
