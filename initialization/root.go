package initialization

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/go-cli/utils"
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

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Init Golang project")
	// nolint: lll
	cmd.Long = t.T("Initialize the Golang project and create default configuration files.")

	cmd.Flags().StringP("license", "", "MIT", t.T("license: APACHE2, BSD3, MIT"))
	cmd.AddCommand(ghCmd(&cobra.Command{Use: "github", Aliases: []string{"g"}}))
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
	code := lo.Must1(cmd.Flags().GetString("license"))
	env.License = strings.ToUpper(code)

	lo.Must0(utils.Walk(_static, _staticPath, func(path string, entry fs.DirEntry) error {
		file, data := readStatic(dir, filepath.Join(path, entry.Name()), env)
		if oss.Exist(file) {
			return nil
		}

		parent := filepath.Dir(file)

		logs.I.Println(t.T("init file %s", file))
		_ = os.MkdirAll(parent, oss.DefaultDirFileMod)

		return os.WriteFile(file, data, oss.DefaultFileMode)
	}))

	license := filepath.Join(dir, "LICENSE")

	if oss.Exist(license) {
		return
	}

	file := lo.Must1(os.Create(license))
	defer file.Close()

	lo.Must1(file.Write(env.Bytes(_licenses, filepath.Join("license", strings.ToLower(code)+".tpl"))))
	logs.I.Println(t.T("init file %s", license))
}

func readStatic(dir, name string, env *tpl.Env) (string, []byte) {
	return readFile(env, _static, dir, name)
}
