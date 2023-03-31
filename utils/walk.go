package utils

import (
	"embed"
	"io/fs"
	"path/filepath"
)

func Walk(files embed.FS, path string, yield func(string, fs.DirEntry) error) error {
	entries, err := files.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		var err error

		if entry.IsDir() {
			err = Walk(files, filepath.Join(path, entry.Name()), yield)
		} else {
			err = yield(path, entry)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
