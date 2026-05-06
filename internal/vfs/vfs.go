package vfs

import (
	"context"
	"io"
	"strings"

	"github.com/Utahamo/nasha/internal/driver"
)

type MountPoint struct {
	Name   string
	Path   string
	Driver driver.StorageDriver
}

type VFS struct {
	mounts []*MountPoint
}

func New() *VFS {
	return &VFS{}
}

func (v *VFS) Mount(mp *MountPoint) {
	v.mounts = append(v.mounts, mp)
}

func (v *VFS) resolve(vpath string) (driver.StorageDriver, string) {
	vpath = "/" + strings.TrimLeft(vpath, "/")
	var best *MountPoint
	for i := range v.mounts {
		mp := v.mounts[i]
		if strings.HasPrefix(vpath, mp.Path) && (best == nil || len(mp.Path) > len(best.Path)) {
			p := mp
			best = p
		}
	}
	if best == nil {
		return nil, ""
	}
	rel := strings.TrimPrefix(vpath, best.Path)
	rel = "/" + strings.TrimLeft(rel, "/")
	return best.Driver, rel
}

func (v *VFS) List(ctx context.Context, vpath string) ([]driver.FileInfo, error) {
	d, rel := v.resolve(vpath)
	if d == nil {
		return nil, errNoMount(vpath)
	}
	return d.List(ctx, rel)
}

func (v *VFS) Read(ctx context.Context, vpath string) (io.ReadCloser, error) {
	d, rel := v.resolve(vpath)
	if d == nil {
		return nil, errNoMount(vpath)
	}
	return d.Read(ctx, rel)
}

func (v *VFS) Write(ctx context.Context, vpath string, r io.Reader) error {
	d, rel := v.resolve(vpath)
	if d == nil {
		return errNoMount(vpath)
	}
	return d.Write(ctx, rel, r)
}

func (v *VFS) Delete(ctx context.Context, vpath string) error {
	d, rel := v.resolve(vpath)
	if d == nil {
		return errNoMount(vpath)
	}
	return d.Delete(ctx, rel)
}

func (v *VFS) MakeDir(ctx context.Context, vpath string) error {
	d, rel := v.resolve(vpath)
	if d == nil {
		return errNoMount(vpath)
	}
	return d.MakeDir(ctx, rel)
}

func (v *VFS) Rename(ctx context.Context, src, dst string) error {
	d, relSrc := v.resolve(src)
	if d == nil {
		return errNoMount(src)
	}
	_, relDst := v.resolve(dst)
	return d.Rename(ctx, relSrc, relDst)
}

func (v *VFS) Stat(ctx context.Context, vpath string) (driver.FileInfo, error) {
	d, rel := v.resolve(vpath)
	if d == nil {
		return driver.FileInfo{}, errNoMount(vpath)
	}
	return d.Stat(ctx, rel)
}

func errNoMount(vpath string) error {
	return &MountError{Path: vpath}
}

type MountError struct {
	Path string
}

func (e *MountError) Error() string {
	return "no mount found for path: " + e.Path
}
