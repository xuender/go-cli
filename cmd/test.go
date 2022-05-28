package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuender/go-scaffold/utils"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

// nolint
func init() {
	root := getRoot()
	testCmd := &cobra.Command{
		Use:   "test",
		Short: Printer.Sprintf("test short"),
		Long:  Printer.Sprintf("test long"),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				Printer.Printf("test missing file")

				return
			}

			var force bool

			if _, err := cmd.Flags().GetBool("force"); err == nil {
				force = true
			}

			for _, arg := range args {
				abs := base.Must1(filepath.Abs(arg))
				if !oss.Exist(abs) {
					panic(Printer.Sprintf("test %s not exist", arg))
				}

				name := filepath.Base(abs)
				if strings.HasSuffix(name, "_test.go") {
					panic(Printer.Sprintf("test not gofile"))
				}

				ext := filepath.Ext(name)
				if ext != ".go" {
					panic(Printer.Sprintf("test not gofile"))
				}

				dir := filepath.Dir(abs)
				out := filepath.Join(dir, name[:len(name)-3]+"_test.go")

				if !force && oss.Exist(out) {
					panic(Printer.Sprintf("test exist"))
				}

				pack, funcs := utils.PackageAndFuncs(abs)
				logs.Debugw("test", "funcs", funcs, "package", pack)

				var buffer bytes.Buffer

				buffer.WriteString("package " + pack + "_test\n\nimport \"testing\"\n")
				for _, fun := range funcs {
					buffer.WriteString("\nfunc Test" + fun + "(t *testing.T) {\n\tt.Parallel()\n\t// TODO\n}\n")
				}

				base.Must(os.WriteFile(out, buffer.Bytes(), oss.DefaultFileMode))
			}
		},
	}

	testCmd.Flags().BoolP("force", "f", false, Printer.Sprintf("test force"))
	root.AddCommand(testCmd)
}
