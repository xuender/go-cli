package font

import (
	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.PersistentFlags().BoolP("all", "a", false, t.T("all characters"))
	cmd.Short = t.T("font tools")
	cmd.Long = t.T("OpenType font subsetter and optimizer.")
	cmd.Example = t.T("  # Font tools\n  go-cli font han")

	cmd.AddCommand(chineseCmd(&cobra.Command{Use: "chinese", Aliases: []string{"c"}}))
	cmd.AddCommand(subsetCmd(&cobra.Command{Use: "subset", Aliases: []string{"s"}}))

	return cmd
}
