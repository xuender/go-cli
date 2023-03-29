package initialization

import (
	"embed"
	"io"
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

//go:embed static
var _static embed.FS

//go:embed license
var _licenses embed.FS

const (
	_staticPath = "static"
	_init       = "init"
)

func SubCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("init Golang project")
	cmd.Long = t.T("init Golang project\n\nfiles:\n  .editorconfig\n  .gitignore\n  .golangci.toml\n  go.mod\n  LICENSE\n  Makefile\n  README.md")

	cmd.Flags().StringP("license", "", "MIT", t.T("license: apache2, bsd3, mit"))

	cmd.Run = run

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	logs.D.Println(t.T("init dir: %s", dir))

	env := tpl.NewEnvByDir(dir)

	for _, entry := range lo.Must1(_static.ReadDir(_staticPath)) {
		file, data := readStatic(dir, entry.Name(), env)

		if oss.Exist(file) {
			continue
		}

		logs.I.Println(t.T("init file %s", file))
		lo.Must0(os.WriteFile(file, data, oss.DefaultFileMode))
	}

	license := filepath.Join(dir, "LICENSE")

	if oss.Exist(license) {
		return
	}

	code := lo.Must1(cmd.Flags().GetString("license"))

	file := lo.Must1(os.Create(license))
	defer file.Close()

	lo.Must1(file.Write(env.Bytes(_licenses, filepath.Join("license", strings.ToLower(code)+".tpl"))))
	logs.I.Println(t.T("init file %s", license))
}

func readStatic(dir, name string, env *tpl.Env) (string, []byte) {
	out := name
	if strings.HasPrefix(name, "d_") {
		out = "." + name[2:]
	}

	tpl := false

	if strings.HasSuffix(out, ".tpl") {
		out = out[:len(out)-4]
		tpl = true
	}

	file := lo.Must1(filepath.Abs(filepath.Join(dir, out)))

	if tpl {
		return file, env.Bytes(_static, filepath.Join(_staticPath, name))
	}

	reader := lo.Must1(_static.Open(filepath.Join(_staticPath, name)))
	defer reader.Close()

	return file, lo.Must1(io.ReadAll(reader))
}
