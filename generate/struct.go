package generate

import (
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

func structCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Generate struct")
	cmd.Long = t.T("Generate struct and new function.")
	cmd.Example = t.T("  # Create struct\n  go-cli g s Book")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)

		for _, arg := range args {
			env := tpl.NewEnvByGo(arg)

			if output != "" {
				env.Path = output
			}

			createStruct(env)
		}
	}

	return cmd
}

func createStruct(env *tpl.Env) {
	logs.D.Println(t.T("create struct: %s", env.Name))

	var file *os.File
	defer file.Close()

	if oss.Exist(env.Path) {
		pkg, names := utils.PackageAndStructs(env.Path)
		env.Package = pkg

		if lo.Contains(names, env.Name) {
			logs.W.Println(t.T("exist: %s", env.Name))

			return
		}

		file = utils.AppendFile(env.Path)
	} else {
		file = utils.CreateFile(env.Path)

		lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, "package.tpl"))))
	}

	lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, "struct.tpl"))))
}
