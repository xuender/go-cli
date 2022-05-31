package cmd

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"gitee.com/xuender/oils/logs"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/utils"
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
		Example: "  go-cli test utils\n  go-cli test utils/parse.go",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New(Printer.Sprintf("test missing file"))
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmdInit(cmd)

			for _, arg := range args {
				createTest(arg)
			}
		},
	}

	root.AddCommand(testCmd)
}

func createTest(arg string) {
	logs.Debug(arg)

	abs := base.Must1(filepath.Abs(arg))

	if !oss.Exist(abs) {
		Printer.Printf("test %s not exist", arg)

		return
	}

	file := base.Must1(os.Stat(abs))

	if file.IsDir() {
		for _, dir := range base.Must1(os.ReadDir(abs)) {
			createTest(filepath.Join(arg, dir.Name()))
		}

		return
	}

	name := filepath.Base(abs)
	if strings.HasSuffix(name, "_test.go") {
		Printer.Printf("test not gofile %s", abs)

		return
	}

	if ext := filepath.Ext(name); ext != ".go" {
		Printer.Printf("test not gofile %s", abs)

		return
	}

	addTests(abs, name)
}

func addTests(abs string, name string) {
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
			Printer.Printf("test %s add %s", out, name)

			count++
		}
	}

	if count == 0 {
		Printer.Printf("test no add %s", out)

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
