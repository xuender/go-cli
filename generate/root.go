package generate

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

//go:embed static
var _static embed.FS

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("generate struct, test, example")
	cmd.Long = t.T("generate:\n  example(e)\n  struct(s)\n  test(t}")

	cmd.PersistentFlags().StringP(_output, "o", "", t.T("output file"))
	cmd.AddCommand(structCmd(&cobra.Command{Use: "struct", Aliases: []string{"s"}}))
	cmd.AddCommand(testCmd(&cobra.Command{Use: "test", Aliases: []string{"t"}}))
	cmd.AddCommand(exampleCmd(&cobra.Command{Use: "example", Aliases: []string{"e"}}))
	cmd.AddCommand(protoCmd(&cobra.Command{Use: "proto", Aliases: []string{"p"}}))
	cmd.AddCommand(cmdCmd(&cobra.Command{Use: "cmd", Aliases: []string{"c"}}))

	return cmd
}
