package generate

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
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

			createGo(env, "struct")
		}
	}

	return cmd
}
