package generate

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func exampleCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Generate example")
	cmd.Long = t.T("Generate test examples for exposed functions in file or directory.")
	cmd.Example = t.T("  # Create example\n  go-cli g e pkg/source.go")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)

		for _, arg := range args {
			createExamples(arg, output)
		}
	}

	return cmd
}

func createExamples(path, output string) {
	if oss.IsDir(path) {
		for _, dir := range lo.Must1(os.ReadDir(path)) {
			createExamples(filepath.Join(path, dir.Name()), output)
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

	createTest(env, "_example_test.go", "example.tpl", "example_func.tpl", "Example")
}
