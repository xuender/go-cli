package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func cmdCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(_type, "t", "", t.T("cobra or flag"))
	cmd.Short = t.T("generate cmd")
	cmd.Long = t.T("generate cmd\n  support for cobra and flag.")
	// nolint: lll
	cmd.Example = t.T("  # create cmd\n  go-cli g c cmd\n  # create cobra\n  go-cli g c cmd -t cobra")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)
		typeCode := lo.Must1(cmd.Flags().GetString(_type))

		if typeCode == "" {
			if file, err := Parse("cmd", "root.go"); err == nil && UseCobra(file) {
				typeCode = _cobra
			}
		}

		if typeCode != _cobra {
			typeCode = "flag"
		}

		for _, arg := range args {
			env := tpl.NewEnvByFile(arg, ".go")

			if output == "" {
				if typeCode == _cobra {
					env.Path = filepath.Join("cmd", env.Path)
				} else {
					base := filepath.Base(env.Path)
					env.Path = filepath.Join("cmd", base[:len(base)-3], "main.go")
				}
			} else {
				env.Path = output
			}

			createCmd(env, typeCode)
		}
	}

	return cmd
}

func createCmd(env *tpl.Env, typeCode string) {
	if oss.Exist(env.Path) {
		logs.W.Println(t.T("exist: %s", env.Path))

		return
	}

	logs.D.Println(t.T("create cmd: %s", env.Name))

	var file *os.File
	defer file.Close()

	file = CreateFile(env.Path)

	lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, fmt.Sprintf("cmd_%s.tpl", typeCode)))))
}

// UseCobra 判断是否使用 spf13/cobra.
func UseCobra(root *ast.File) bool {
	for _, imp := range root.Imports {
		if strings.Contains(imp.Path.Value, "spf13/cobra") {
			return true
		}
	}

	return false
}

// Parse Go源程序解析.
func Parse(files ...string) (*ast.File, error) {
	filename, err := filepath.Abs(filepath.Join(files...))
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()

	return parser.ParseFile(fset, filename, nil, parser.ParseComments)
}
