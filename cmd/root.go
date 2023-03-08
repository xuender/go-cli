package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/oils/i18n"
	"github.com/xuender/oils/oss"
)

//go:embed locales
var locales embed.FS

// nolint
var Printer = i18n.GetPrinter()

// nolint
var rootCmd *cobra.Command

func getRoot() *cobra.Command {
	if rootCmd == nil {
		lo.Must0(i18n.Load(locales))

		mod := oss.GetMod("cmd")
		rootCmd = &cobra.Command{
			Use:     "go-cli",
			Short:   Printer.Sprintf("root short"),
			Long:    Printer.Sprintf("root long"),
			Version: fmt.Sprintf("%s [%s]", mod.Version, mod.Sum),
			Example: "  go-cli init\n  go-cli g\n  go-cli g cmd\n  go-cli g service pkg/service\n  go-cli test pkg/service.go",
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
func init() {
	root := getRoot()
	root.PersistentFlags().BoolP("debug", "d", false, Printer.Sprintf("root debug"))
	root.PersistentFlags().StringP("language", "l", "", Printer.Sprintf("root language"))
}
