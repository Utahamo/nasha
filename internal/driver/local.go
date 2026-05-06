package driver

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type LocalDriver struct {
	Root string
}

func (d *LocalDriver) resolve(vpath string) string {
	return filepath.Join(d.Root, filepath.Clean("/"+vpath))
}

func (d *LocalDriver) List(_ context.Context, vpath string) ([]FileInfo, error) {
	dir := d.resolve(vpath)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	files := make([]FileInfo, 0, len(entries))
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}
		files = append(files, FileInfo{
			Name:    e.Name(),
			Size:    info.Size(),
			IsDir:   e.IsDir(),
			ModTime: info.ModTime(),
			Path:    normalJoin(vpath, e.Name()),
		})
	}
	return files, nil
}

func (d *LocalDriver) Read(_ context.Context, vpath string) (io.ReadCloser, error) {
	return os.Open(d.resolve(vpath))
}

func (d *LocalDriver) Write(_ context.Context, vpath string, r io.Reader) error {
	full := d.resolve(vpath)
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}
	f, err := os.Create(full)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	cerr := f.Close()
	if err != nil {
		return err
	}
	return cerr
}

func (d *LocalDriver) Delete(_ context.Context, vpath string) error {
	return os.RemoveAll(d.resolve(vpath))
}

func (d *LocalDriver) MakeDir(_ context.Context, vpath string) error {
	return os.MkdirAll(d.resolve(vpath), 0755)
}

func (d *LocalDriver) Rename(_ context.Context, src, dst string) error {
	return os.Rename(d.resolve(src), d.resolve(dst))
}

func (d *LocalDriver) Stat(_ context.Context, vpath string) (FileInfo, error) {
	info, err := os.Stat(d.resolve(vpath))
	if err != nil {
		return FileInfo{}, err
	}
	return FileInfo{
		Name:    info.Name(),
		Size:    info.Size(),
		IsDir:   info.IsDir(),
		ModTime: info.ModTime(),
		Path:    vpath,
	}, nil
}

// normalJoin joins two path segments with a forward slash,
// ensuring the result is a clean path with / separators regardless of OS.
func normalJoin(a, b string) string {
	a = "/" + strings.TrimLeft(a, "/")
	if a == "/" {
		return "/" + b
	}
	return strings.TrimRight(a, "/") + "/" + b
}
