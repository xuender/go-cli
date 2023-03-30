package generate

import (
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func TestCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("generate test")
	cmd.Long = t.T("generate test")
	cmd.Example = t.T("  # create test\n  go-cli g t pkg/source.go")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)

		for _, arg := range args {
			env := tpl.NewEnvByGo(arg)

			if output != "" {
				env.Test = output
			}

			createTest(env, "_test.go", "test.tpl", "test_func.tpl")
		}
	}

	return cmd
}

func createTest(env *tpl.Env, ext, headFile, funcFile string) {
	logs.D.Println(t.T("create test: %s", env.Name))

	if !oss.Exist(env.Path) {
		logs.E.Println(t.T("not found: %s", env.Path))

		return
	}

	if env.Test == "" {
		env.Test = env.Path[:len(env.Path)-3] + ext
	}

	pkg, funcs := tpl.PackageAndFuncs(env.Path)
	tests := []string{}

	env.Package = pkg

	var file *os.File
	defer file.Close()

	if oss.Exist(env.Test) {
		file = AppendFile(env.Test)
		_, tests = tpl.PackageAndFuncs(env.Test)
	} else {
		file = CreateFile(env.Test)
		lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, headFile))))
	}

	for _, name := range funcs {
		if lo.Contains(tests, "Test"+name) {
			continue
		}

		logs.D.Println(t.T("create: %s %s", env.Test, name))
		env.Name = name

		lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, funcFile))))
	}
}
