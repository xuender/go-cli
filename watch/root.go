package watch

import (
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/los"
	"github.com/youthlin/t"
)

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Watch and Run")
	cmd.Long = t.T("Watch to the directory and run a command. If the directory is modified, restart the command.")
	cmd.Example = t.T("  # Watch dir\n  go-cli watch [command]")

	cmd.Flags().StringP("path", "p", "", t.T("watch path, default is current directory"))
	cmd.Run = run

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		lo.Must0(cmd.Usage())

		return
	}

	logs.D.Println(strings.Join(args, " "))

	if strings.Contains(args[0], " ") {
		args = append(los.SplitStr(args[0], ' ', '\t'), args[1:]...)
	}

	watch := NewService()
	defer watch.Close()

	if path, err := cmd.Flags().GetString("path"); err == nil && path != "" {
		watch.Add(path)
	} else {
		watch.Add(".")
	}

	watch.Run(args[0], args[1:])
}
