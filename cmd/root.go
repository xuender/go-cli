package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/font"
	"github.com/xuender/go-cli/generate"
	"github.com/xuender/go-cli/initialization"
	"github.com/xuender/go-cli/structs"
	"github.com/xuender/go-cli/version"
	"github.com/xuender/go-cli/watch"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "go-cli",
		Short: t.T("CLI tool for Golang"),
		Long: t.T("CLI tool for Golang\n\n  Generate structures, tests, examples, initialize projects, etc.\n\n\t\t%s",
			oss.BuildTime),
		Version: oss.Version,
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringP("language", "l", t.Global().Locale(), t.T("select language: en, zh"))
	rootCmd.PersistentFlags().BoolP("debug", "d", false, t.T("Debug mode, display debug log"))
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if debug, err := cmd.Flags().GetBool("debug"); err == nil && debug {
			logs.D.Println(t.T("set debug: %v", debug))
			logs.SetLevel(logs.Debug)
		} else {
			logs.SetLevel(logs.Info)
		}

		if lang, err := cmd.Flags().GetString("language"); err == nil {
			logs.D.Println(t.T("set language: %s", lang))
			t.SetLocale(lang)
		}
	}
	rootCmd.AddCommand(initialization.NewCmd(&cobra.Command{Use: "init", Aliases: []string{"i", "initialization"}}))
	rootCmd.AddCommand(generate.NewCmd(&cobra.Command{Use: "generate", Aliases: []string{"g", "creaet"}}))
	rootCmd.AddCommand(structs.NewCmd(&cobra.Command{Use: "struct", Aliases: []string{"s"}}))
	rootCmd.AddCommand(watch.NewCmd(&cobra.Command{Use: "watch", Aliases: []string{"w"}}))
	rootCmd.AddCommand(TemplateCmd(&cobra.Command{Use: "template", Aliases: []string{"t"}}))
	rootCmd.AddCommand(version.NewCmd(&cobra.Command{Use: "version", Aliases: []string{"v", "ver"}}))
	rootCmd.AddCommand(font.NewCmd(&cobra.Command{Use: "font", Aliases: []string{"f"}}))

	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
