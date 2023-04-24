package structs

import (
	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Struct related")
	cmd.Long = t.T("Struct related commands.")

	cmd.AddCommand(newStruct(&cobra.Command{Use: "new", Aliases: []string{"n"}}))
	cmd.AddCommand(convertStruct(&cobra.Command{Use: "convert", Aliases: []string{"c"}}))

	return cmd
}
