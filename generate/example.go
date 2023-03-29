package generate

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/youthlin/t"
)

func ExampleCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("generate example")
	cmd.Long = t.T("generate example")
	cmd.Example = t.T("  # create example\n  go-cli g e pkg/source.go")
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

			CreateTest(env, "_example_test.go", "example.tpl", "example_func.tpl")
		}
	}

	return cmd
}
