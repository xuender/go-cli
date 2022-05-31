package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"
	"strings"
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
				if name := getName(node); name != "" {
					funcs = append(funcs, name)
				}
			}
		case *ast.File:
			pack = node.Name.Name
		}

		return true
	})

	return pack, funcs
}

func getName(node *ast.FuncDecl) string {
	if node.Recv == nil {
		return node.Name.Name
	}

	for _, field := range node.Recv.List {
		switch elem := field.Type.(type) {
		case *ast.StarExpr:
			if ident, ok := elem.X.(*ast.Ident); ok && unicode.IsUpper([]rune(ident.Name)[0]) {
				return ident.Name + "_" + node.Name.Name
			}
		case *ast.Ident:
			if unicode.IsUpper([]rune(elem.Name)[0]) {
				return elem.Name + "_" + node.Name.Name
			}
		}
	}

	return ""
}

func GitURL(data []byte) string {
	reg := regexp.MustCompile(`[A-Za-z0-9_\-./:]+\.git`)
	ret := string(reg.Find(data))
	ret = strings.TrimSuffix(ret, ".git")
	ret = strings.TrimPrefix(ret, "https://")
	ret = strings.TrimPrefix(ret, "http://")

	return strings.Replace(ret, ":", "/", 1)
}
