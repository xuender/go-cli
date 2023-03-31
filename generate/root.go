package generate

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

//go:embed static
var _static embed.FS

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Generate source code")
	// nolint: lll
	cmd.Long = t.T("Generate source code including commands, tests, examples, structures, protobuf, etc.")

	cmd.PersistentFlags().StringP(_output, "o", "", t.T("Output file"))
	cmd.AddCommand(structCmd(&cobra.Command{Use: "struct", Aliases: []string{"s"}}))
	cmd.AddCommand(testCmd(&cobra.Command{Use: "test", Aliases: []string{"t"}}))
	cmd.AddCommand(exampleCmd(&cobra.Command{Use: "example", Aliases: []string{"e"}}))
	cmd.AddCommand(protoCmd(&cobra.Command{Use: "proto", Aliases: []string{"p"}}))
	cmd.AddCommand(cmdCmd(&cobra.Command{Use: "cmd", Aliases: []string{"c"}}))

	return cmd
}
