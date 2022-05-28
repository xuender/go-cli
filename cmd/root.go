/*
Copyright © 2022 Anicca.cn <xuender@139.com>

*/
package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/i18n"
	"golang.org/x/text/language"
)

//go:embed locales
var locales embed.FS

// nolint
var Printer = i18n.GetPrinter(language.SimplifiedChinese, language.English)

// nolint
var rootCmd *cobra.Command

func getRoot() *cobra.Command {
	if rootCmd == nil {
		base.Must(i18n.Load(locales))

		rootCmd = &cobra.Command{
			Use:   "gos",
			Short: Printer.Sprintf("root short"),
			Long:  Printer.Sprintf("root long"),
			// Run: func(cmd *cobra.Command, args []string) { },
		}

		rootCmd.CompletionOptions.DisableDefaultCmd = true
	}

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// nolint
// func init() {
// 	getRoot().Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
