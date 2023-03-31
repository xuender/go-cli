package generate

import (
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

func testCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Generate test")
	cmd.Long = t.T("Generating unit tests for exposed functions in file or directory.")
	cmd.Example = t.T("  # Create test\n  go-cli g t pkg/source.go\n  # Create path\n  go-cli g t pkg")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)

		for _, arg := range args {
			createTests(arg, output)
		}
	}

	return cmd
}

func createTests(path, output string) {
	if oss.IsDir(path) {
		for _, dir := range lo.Must1(os.ReadDir(path)) {
			createTests(filepath.Join(path, dir.Name()), output)
		}

		return
	}

	if strings.HasSuffix(path, "_test.go") || !strings.HasSuffix(path, ".go") {
		return
	}

	env := tpl.NewEnvByGo(path)

	if output != "" {
		env.Test = output
	}

	createTest(env, "_test.go", "test.tpl", "test_func.tpl")
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

	pkg, funcs := utils.PackageAndFuncs(env.Path)
	if len(funcs) == 0 {
		return
	}

	tests := []string{}

	env.Package = pkg

	var file *os.File
	defer file.Close()

	if oss.Exist(env.Test) {
		file = utils.AppendFile(env.Test)
		_, tests = utils.PackageAndFuncs(env.Test)
	} else {
		file = utils.CreateFile(env.Test)
		lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, headFile))))
	}

	for _, name := range funcs {
		if lo.Contains(tests, "Test"+name) {
			continue
		}

		logs.I.Println(t.T("create: %s %s", env.Test, name))
		env.Name = name

		lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, funcFile))))
	}
}
