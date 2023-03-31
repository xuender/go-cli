package tpl

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

type Dir interface {
	ReadDir(string) ([]fs.DirEntry, error)
	Open(string) (fs.File, error)
}

type DirEntry struct {
	path string
}

func NewDirEntry(path string) *DirEntry {
	return &DirEntry{path}
}

func (p *DirEntry) ReadDir(path string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(p.path, path))
}

func (p *DirEntry) Open(path string) (fs.File, error) {
	logs.D.Println("open", p.path, path)

	return os.OpenFile(filepath.Join(p.path, path), os.O_RDONLY, oss.DefaultFileMode)
}

func (p *DirEntry) String() string {
	return fmt.Sprintf("DirEntry: %s", p.path)
}
