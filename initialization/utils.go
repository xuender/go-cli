package initialization

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func getRun(path string, files embed.FS) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		logs.D.Println(t.T("init dir: %s", dir))

		env := tpl.NewEnvByDir(dir)

		lo.Must0(tpl.Walk(files, path, func(path string, entry fs.DirEntry) error {
			file, data := readFile(env, files, dir, filepath.Join(path, entry.Name()))
			if oss.Exist(file) {
				return nil
			}

			logs.I.Println(t.T("init file %s", file))

			parent := filepath.Dir(file)

			_ = os.MkdirAll(parent, oss.DefaultDirFileMod)

			return os.WriteFile(file, data, oss.DefaultFileMode)
		}))
	}
}

func readFile(env *tpl.Env, dir tpl.Dir, target, path string) (string, []byte) {
	logs.D.Println("target:", target, "path:", path, "dir", dir)

	sep := string(os.PathSeparator)
	out := strings.Join(lo.Map(strings.Split(path, sep), func(name string, _ int) string {
		if strings.HasPrefix(name, "d_") {
			return "." + name[2:]
		}

		return name
	}), sep)

	tpl := false

	if strings.HasSuffix(out, ".tpl") {
		out = out[:len(out)-4]
		tpl = true
	}

	if index := strings.Index(out, sep); index > 0 {
		out = out[index:]
	}

	file := lo.Must1(filepath.Abs(filepath.Join(target, out)))

	if tpl {
		logs.D.Println("file:", file, "reader:", path)

		return file, env.Bytes(dir, path)
	}

	reader := lo.Must1(dir.Open(path))
	defer reader.Close()

	logs.D.Println("file:", file, "reader:", path)

	return file, lo.Must1(io.ReadAll(reader))
}
