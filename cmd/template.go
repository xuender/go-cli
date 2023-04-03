package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/generate"
	"github.com/xuender/go-cli/initialization"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/logs"
	"github.com/youthlin/t"
)

func TemplateCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Create templates")
	cmd.Long = t.T("Create templates to %s", tpl.ConfigPath)
	cmd.Run = func(cmd *cobra.Command, args []string) {
		logs.D.Println(t.T("create %s templates.", "generate"))
		tpl.WriteTemplate(generate.GenStatic)
		logs.D.Println(t.T("create %s templates.", "init"))
		tpl.WriteTemplate(initialization.InitStatic)
		logs.D.Println(t.T("create %s templates.", "github"))
		tpl.WriteTemplate(initialization.GithubStatic)
		logs.D.Println(t.T("create %s templates.", "gitee"))
		tpl.WriteTemplate(initialization.GiteeStatic)
	}

	return cmd
}
