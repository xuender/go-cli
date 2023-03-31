package tpl

import (
	"io/fs"
	"path/filepath"
)

func Walk(dir Dir, path string, yield func(string, fs.DirEntry) error) error {
	entries, err := dir.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		var err error

		if entry.IsDir() {
			err = Walk(dir, filepath.Join(path, entry.Name()), yield)
		} else {
			err = yield(path, entry)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
