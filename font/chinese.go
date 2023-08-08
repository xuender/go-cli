package font

import (
	"fmt"
	"os"
	"unicode"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/ordered"
	"github.com/xuender/kit/oss"
	"github.com/xuender/kit/set"
	"github.com/youthlin/t"
)

func chineseCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Extracting Chinese")
	cmd.Long = t.T("Extracting Chinese Characters from file.")
	cmd.Example = t.T("  # Extracting Chinese\n  go-cli font chinese zh_CN.po")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		all, _ := cmd.Flags().GetBool("all")

		fmt.Fprintln(os.Stdout, load(all, args...))
	}

	return cmd
}

func load(all bool, args ...string) string {
	chinese := set.NewSet[rune]()

	for _, arg := range args {
		logs.D.Println(t.T("load: %s", arg))

		for _, code := range string(lo.Must1(os.ReadFile(lo.Must1(oss.Abs(arg))))) {
			if unicode.IsSpace(code) {
				continue
			}

			if all || unicode.Is(unicode.Scripts["Han"], code) {
				chinese.Add(code)
			}
		}
	}

	slice := chinese.Slice()

	ordered.Sort(slice)

	return string(slice)
}
