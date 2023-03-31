package initialization

import (
	"embed"

	"github.com/spf13/cobra"
	"github.com/youthlin/t"
)

//go:embed github
var _github embed.FS

func ghCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Init github config")
	cmd.Long = t.T("Initialize the github configuration files.")
	cmd.Run = getRun("github", _github)

	return cmd
}
