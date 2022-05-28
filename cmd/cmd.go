/*
Copyright © 2022 Anicca.cn <xuender@139.com>

*/
package cmd

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xuender/go-scaffold/utils"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

//go:embed cmd.txt
var defaultCmd []byte

// nolint
func init() {
	root := getRoot()
	cmdCmd := &cobra.Command{
		Use:   "cmd",
		Short: Printer.Sprintf("cmd short"),
		Long:  Printer.Sprintf("cmd long"),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				Printer.Printf("cmd missing command")

				return
			}

			file, err := utils.Parse("cmd", "root.go")
			useCobra := err == nil && utils.UseCobra(file)

			if cobra, err := cmd.Flags().GetBool("cobra"); err == nil && cobra {
				useCobra = true
			}

			if flag, err := cmd.Flags().GetBool("flag"); err == nil && flag {
				useCobra = false
			}

			for _, arg := range args {
				logs.Debugw("创建", "arg", arg)
				if useCobra {
					if oss.Exist("cmd", arg+".go") {
						panic(Printer.Sprintf("cmd %s exist", arg))
					}

					base.Must(oss.Exec("cobra-cli", "add", arg))

					continue
				}

				if oss.Exist("cmd", arg) {
					panic(Printer.Sprintf("cmd %s exist", arg))
				}

				base.Must(os.MkdirAll(filepath.Join("cmd", arg), oss.DefaultDirFileMod))
				base.Must(os.WriteFile(
					filepath.Join("cmd", arg, "main.go"),
					bytes.ReplaceAll(defaultCmd, []byte("NAME"), []byte(arg)),
					oss.DefaultFileMode,
				))
			}
		},
	}

	cmdCmd.Flags().BoolP("cobra", "c", false, Printer.Sprintf("cmd cobra"))
	cmdCmd.Flags().BoolP("flag", "f", false, Printer.Sprintf("cmd flag"))
	root.AddCommand(cmdCmd)
}
