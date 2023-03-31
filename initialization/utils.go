package initialization

import (
	"embed"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/logs"
)

func readFile(env *tpl.Env, static embed.FS, target, path string) (string, []byte) {
	logs.D.Println("target:", target, "path:", path)

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
		out = out[:index]
	}

	file := lo.Must1(filepath.Abs(filepath.Join(target, out)))

	if tpl {
		return file, env.Bytes(static, path)
	}

	reader := lo.Must1(static.Open(path))
	defer reader.Close()

	logs.D.Println("file:", file, "reader:", path)

	return file, lo.Must1(io.ReadAll(reader))
}
