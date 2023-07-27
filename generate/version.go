package generate

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/go-cli/utils"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func versionCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Generate version")
	cmd.Long = t.T("Generate version_gen.go.")
	cmd.Example = t.T("  # Create version_gen.go\n  go-cli g v cmd")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)

		if output == "" {
			output = args[0]
		}

		pack := filepath.Base(output)
		path := output

		if oss.Exist(output) {
			if oss.IsDir(output) {
				pack = filepath.Base(output)
			} else {
				pack = filepath.Dir(output)
			}
		}

		if !strings.HasSuffix(output, ".go") {
			path = filepath.Join(output, "version_gen.go")
		}

		env := &tpl.Env{
			BuildTime: time.Now().Format("2006-01-02 15:04:05"),
			Package:   pack,
			Version:   "development",
		}

		comm := exec.Command("git", "describe", "--tags")
		if ver, err := comm.CombinedOutput(); err == nil {
			env.Version = strings.TrimSpace(string(ver))
		}

		var file *os.File
		defer file.Close()

		file = utils.CreateFile(path)
		lo.Must1(file.Write(env.Bytes(_dir, filepath.Join(_staticPath, "version_gen.tpl"))))
	}

	return cmd
}
