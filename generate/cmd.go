package generate

import (
	"fmt"
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

func cmdCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(_type, "t", "", t.T("cobra or flag"))
	cmd.Short = t.T("generate cmd")
	cmd.Long = t.T("generate cmd\n  support for cobra and flag.")
	// nolint: lll
	cmd.Example = t.T("  # create cmd\n  go-cli g c cmd\n  # create cobra\n  go-cli g c cmd -t cobra")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)
		typeCode := lo.Must1(cmd.Flags().GetString(_type))

		if typeCode == "" {
			if file, err := utils.Parse("cmd", "root.go"); err == nil && utils.UseCobra(file) {
				typeCode = _cobra
			}
		}

		if typeCode != _cobra {
			typeCode = "flag"
		}

		for _, arg := range args {
			env := tpl.NewEnvByFile(arg, ".go")

			if output == "" {
				if typeCode == _cobra {
					env.Path = filepath.Join("cmd", env.Path)
				} else {
					base := filepath.Base(env.Path)
					env.Path = filepath.Join("cmd", base[:len(base)-3], "main.go")
				}
			} else {
				env.Path = output
			}

			createCmd(env, typeCode)
		}
	}

	return cmd
}

func createCmd(env *tpl.Env, typeCode string) {
	if oss.Exist(env.Path) {
		logs.W.Println(t.T("exist: %s", env.Path))

		return
	}

	logs.D.Println(t.T("create cmd: %s", env.Name))

	var file *os.File
	defer file.Close()

	file = utils.CreateFile(env.Path)

	lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, fmt.Sprintf("cmd_%s.tpl", typeCode)))))
}
