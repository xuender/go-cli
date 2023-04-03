package initialization

import (
	"bytes"
	"embed"
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

//go:embed init
var InitStatic embed.FS

//go:embed license
var _licenses embed.FS

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Init Golang project")
	cmd.Long = t.T("Initialize the Golang project and create default configuration files.")
	cmd.Example = t.T("  # Init project\n  go-cli init\n  # Init github config\n go-cli init github")

	cmd.Flags().StringP("license", "", "MIT", t.T("license: APACHE2, BSD3, MIT"))
	cmd.AddCommand(githubCmd(&cobra.Command{Use: "github", Aliases: []string{"gh"}}))
	cmd.AddCommand(geteeCmd(&cobra.Command{Use: "gitee", Aliases: []string{"ge"}}))
	cmd.Run = run

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	code := lo.Must1(cmd.Flags().GetString("license"))
	env := tpl.NewEnvByDir(".")
	env.License = strings.ToUpper(code)

	if len(args) > 0 {
		for _, arg := range args {
			logs.D.Println(t.T("init: %s", arg))

			if dir := filepath.Join(tpl.ConfigPath, arg); !oss.Exist(dir) {
				logs.W.Println(t.T("dir is not exist: %s", dir))

				continue
			}

			initFiles(tpl.NewDirEntry(tpl.ConfigPath), arg, env)
		}

		return
	}

	if dir := filepath.Join(tpl.ConfigPath, "init"); oss.Exist(dir) {
		initFiles(tpl.NewDirEntry(tpl.ConfigPath), "init", env)
	} else {
		initFiles(InitStatic, "init", env)
	}

	license := filepath.Join(".", "LICENSE")

	if oss.Exist(license) {
		return
	}

	file := lo.Must1(os.Create(license))
	defer file.Close()

	lo.Must1(file.Write(env.Bytes(_licenses, filepath.Join("license", strings.ToLower(code)+".tpl"))))
	logs.I.Println(t.T("init file %s", license))
}

func initFiles(dir tpl.Dir, path string, env *tpl.Env) {
	lo.Must0(tpl.Walk(dir, path, func(path string, entry fs.DirEntry) error {
		file, data := readFile(env, dir, ".", filepath.Join(path, entry.Name()))
		if oss.Exist(file) {
			return nil
		}

		parent := filepath.Dir(file)

		logs.I.Println(t.T("init file %s", file))
		_ = os.MkdirAll(parent, oss.DefaultDirFileMod)

		if bytes.HasPrefix(data, []byte("#!")) {
			return os.WriteFile(file, data, oss.DefaultDirFileMod)
		}

		return os.WriteFile(file, data, oss.DefaultFileMode)
	}))
}
