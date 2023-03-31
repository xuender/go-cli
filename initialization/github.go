package initialization

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/go-cli/utils"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

//go:embed github
var _github embed.FS

func ghCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Init github project")
	cmd.Long = t.T("Init github project\n\nCreate:\n  .github/workflows/go.yml")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		logs.D.Println(t.T("init github dir: %s", dir))

		env := tpl.NewEnvByDir(dir)

		lo.Must0(utils.Walk(_github, "github", func(path string, entry fs.DirEntry) error {
			file, data := readGithub(dir, filepath.Join(path, entry.Name()), env)

			if oss.Exist(file) {
				return nil
			}

			logs.I.Println(t.T("init file %s", file))

			parent := filepath.Dir(file)

			_ = os.MkdirAll(parent, oss.DefaultDirFileMod)

			return os.WriteFile(file, data, oss.DefaultFileMode)
		}))
	}

	return cmd
}

func readGithub(dir, name string, env *tpl.Env) (string, []byte) {
	return readFile(env, _github, dir, name)
}
