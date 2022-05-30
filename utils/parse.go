package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"unicode"
)

// Parse Go源程序解析.
func Parse(files ...string) (*ast.File, error) {
	filename, err := filepath.Abs(filepath.Join(files...))
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()

	return parser.ParseFile(fset, filename, nil, parser.ParseComments)
}

// PackageAndFuncs 解析Go源程序包和方法.
func PackageAndFuncs(files ...string) (string, []string) {
	funcs := []string{}

	file, err := Parse(files...)
	if err != nil {
		return "", funcs
	}

	pack := ""

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			if node.Name.Name != "" && unicode.IsUpper([]rune(node.Name.Name)[0]) {
				funcs = append(funcs, node.Name.Name)
			}
		case *ast.File:
			pack = node.Name.Name
		}

		return true
	})

	return pack, funcs
}
