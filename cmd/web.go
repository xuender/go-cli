package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/oils/i18n"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/nets"
)

// nolint
func init() {
	var port int16 = 8080

	root := getRoot()
	webCmd := &cobra.Command{
		Use:     "web",
		Short:   Printer.Sprintf("web short"),
		Long:    Printer.Sprintf("web long"),
		Example: "  go-cli web [path...]",
		Aliases: []string{"w"},
		Run: func(cmd *cobra.Command, args []string) {
			if debug, err := cmd.Flags().GetBool("debug"); err != nil || !debug {
				logs.SetInfoLevel()
			}

			if language, err := cmd.Flags().GetString("language"); err == nil && language != "" {
				Printer = i18n.GetPrinter(i18n.GetTag([]string{language}))
			}

			if portVal, err := cmd.Flags().GetInt16("port"); err == nil {
				port = portVal
			}

			dir := lo.Must1(os.Getwd())

			if len(args) == 1 {
				dir = args[0]
			}

			var files http.FileSystem
			if len(args) < 2 {
				files = http.Dir(dir)
			} else {
				files = nets.Dirs(args)
			}

			logs.Infow("web", "port", port, "paths", args)
			lo.Must0(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(files)))
		},
	}

	webCmd.Flags().Int16P("port", "p", port, Printer.Sprintf("web port"))
	root.AddCommand(webCmd)
}
