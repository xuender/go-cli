package cmd

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuender/go-scaffold/utils"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/oss"
)

// nolint
func init() {
	root := getRoot()
	testCmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"t"},
		Short:   Printer.Sprintf("test short"),
		Long:    Printer.Sprintf("test long"),
		Example: "  go-scaffold test utils.go",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New(Printer.Sprintf("test missing file"))
			}

			for _, arg := range args {
				if !oss.Exist(arg) {
					return errors.New(Printer.Sprintf("test %s not exist", arg))
				}
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				createTest(arg)
			}
		},
	}

	root.AddCommand(testCmd)
}

func createTest(arg string) {
	abs := base.Must1(filepath.Abs(arg))

	name := filepath.Base(abs)
	if strings.HasSuffix(name, "_test.go") {
		panic(Printer.Sprintf("test not gofile"))
	}

	if ext := filepath.Ext(name); ext != ".go" {
		panic(Printer.Sprintf("test not gofile"))
	}

	dir := filepath.Dir(abs)
	out := filepath.Join(dir, name[:len(name)-3]+"_test.go")
	pack, funcs := utils.PackageAndFuncs(abs)
	tests := []string{}

	var (
		exist  bool
		buffer bytes.Buffer
	)

	exist = oss.Exist(out)

	if exist {
		_, tests = utils.PackageAndFuncs(out)
	} else {
		buffer.WriteString("package " + pack + "_test\n\nimport \"testing\"\n")
	}

	count := 0

	for _, fun := range funcs {
		name := "Test" + fun
		if !base.Has(tests, name) {
			buffer.WriteString("\nfunc " + name + "(t *testing.T) {\n\tt.Parallel()\n\t// TODO\n}\n")
			Printer.Printf("test add %s", name)

			count++
		}
	}

	if count == 0 {
		Printer.Printf("test no add %s", arg)

		return
	}

	if exist {
		file := base.Must1(os.OpenFile(out, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm))
		defer file.Close()

		base.Must1(file.Write(buffer.Bytes()))
	} else {
		base.Must(os.WriteFile(out, buffer.Bytes(), oss.DefaultFileMode))
	}
}
