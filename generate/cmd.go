package generate

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

//go:embed static
var _static embed.FS

func SubCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("generate struct, test, example")
	cmd.Long = t.T("generate:\n  example(e)\n  struct(s)\n  test(t}")

	cmd.PersistentFlags().StringP(_output, "o", "", t.T("output file"))
	cmd.AddCommand(StructCmd(&cobra.Command{Use: _struct, Aliases: []string{"s"}}))
	cmd.AddCommand(TestCmd(&cobra.Command{Use: _test, Aliases: []string{"t"}}))
	cmd.AddCommand(ExampleCmd(&cobra.Command{Use: _example, Aliases: []string{"e"}}))
	cmd.AddCommand(ProtoCmd(&cobra.Command{Use: _proto, Aliases: []string{"p"}}))

	return cmd
}
