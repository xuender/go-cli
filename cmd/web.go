/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
)

// nolint
func init() {
	var port int16 = 8080

	root := getRoot()
	webCmd := &cobra.Command{
		Use:     "web",
		Short:   Printer.Sprintf("web short"),
		Long:    Printer.Sprintf("web long"),
		Example: "  go-cli web",
		Aliases: []string{"w"},
		Run: func(cmd *cobra.Command, args []string) {
			cmdInit(cmd)

			if portVal, err := cmd.Flags().GetInt16("port"); err == nil {
				port = portVal
			}

			dir := base.Must1(os.Getwd())

			if len(args) > 0 {
				dir = args[0]
			}

			logs.Debugw("web", "port", port, "path", dir)
			base.Must(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
		},
	}

	webCmd.Flags().Int16P("port", "p", port, Printer.Sprintf("web port"))
	root.AddCommand(webCmd)
}
