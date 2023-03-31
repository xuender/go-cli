package initialization

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

//go:embed gitee
var _gitee embed.FS

func geteeCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Init gitee config")
	cmd.Long = t.T("Initialize the gitee configuration files.")
	cmd.Run = getRun("gitee", _gitee)

	return cmd
}
