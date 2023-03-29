package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/generate"
	"github.com/xuender/go-cli/initialization"
	"github.com/xuender/kit/logs"
	"github.com/youthlin/t"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:     "go-cli",
		Short:   t.T("CLI tool for Golang"),
		Long:    t.T("CLI tool for Golang\n\n  Generate structures, tests, examples, initialize projects, etc."),
		Version: "1.1.10",
	}

	rootCmd.PersistentFlags().StringP("language", "l", t.Global().Locale(), t.T("select language: en, zh"))
	rootCmd.PersistentFlags().BoolP("debug", "d", false, t.T("debug mode, display debug log"))
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
	rootCmd.AddCommand(initialization.SubCmd(&cobra.Command{Use: "init", Aliases: []string{"i", "initialization"}}))
	rootCmd.AddCommand(generate.SubCmd(&cobra.Command{Use: "generate", Aliases: []string{"g", "creaet"}}))

	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
