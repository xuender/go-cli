package generate

import (
	"embed"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

//go:embed gen
var _static embed.FS

// nolint: gochecknoglobals
var _dir tpl.Dir

func NewCmd(cmd *cobra.Command) *cobra.Command {
	if oss.Exist(filepath.Join(tpl.ConfigPath, "gen")) {
		_dir = tpl.NewDirEntry(tpl.ConfigPath)
	} else {
		_dir = _static
	}

	cmd.Short = t.T("Generate source code")
	cmd.Long = t.T("Generate source code including commands, tests, examples, structures, protobuf, etc.")

	cmd.PersistentFlags().StringP(_output, "o", "", t.T("Output file"))
	cmd.AddCommand(structCmd(&cobra.Command{Use: "struct", Aliases: []string{"s"}}))
	cmd.AddCommand(testCmd(&cobra.Command{Use: "test", Aliases: []string{"t"}}))
	cmd.AddCommand(exampleCmd(&cobra.Command{Use: "example", Aliases: []string{"e"}}))
	cmd.AddCommand(protoCmd(&cobra.Command{Use: "proto", Aliases: []string{"p"}}))
	cmd.AddCommand(cmdCmd(&cobra.Command{Use: "cmd", Aliases: []string{"c"}}))
	cmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		tpl.WriteTemplate(cmd, _static)
	}

	return cmd
}
