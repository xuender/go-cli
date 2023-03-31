package tpl

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

func WriteTemplate(cmd *cobra.Command, files embed.FS) {
	if write, err := cmd.Flags().GetBool("write"); err != nil || !write {
		return
	}

	lo.Must0(Walk(files, ".", func(path string, entry fs.DirEntry) error {
		logs.D.Println("template", path, entry.Name())

		file := filepath.Join(lo.Must1(os.UserHomeDir()), ".config", "go-cli", path, entry.Name())
		if oss.Exist(file) {
			return nil
		}

		dir := filepath.Dir(file)
		_ = os.MkdirAll(dir, oss.DefaultDirFileMod)

		return os.WriteFile(file, lo.Must1(files.ReadFile(filepath.Join(path, entry.Name()))), oss.DefaultFileMode)
	}))
}
