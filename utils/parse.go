package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
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

func PackageAndStructs(files ...string) (string, []string) {
	structs := []string{}

	file, err := Parse(files...)
	if err != nil {
		return "", structs
	}

	pack := ""

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSpec:
			if _, isStruct := node.Type.(*ast.StructType); isStruct {
				structs = append(structs, node.Name.Name)
			}
		case *ast.File:
			pack = node.Name.Name
		}

		return true
	})

	return pack, structs
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

func getX(expr ast.Expr) string {
	ret := ""

	switch elem := expr.(type) {
	case *ast.IndexListExpr:
		return getX(elem.X)
	case *ast.Ident:
		ret = elem.Name
	case *ast.IndexExpr:
		ret = getX(elem.X)
	}

	if ret != "" && unicode.IsUpper([]rune(ret)[0]) {
		return ret
	}

	return ""
}

func getName(node *ast.FuncDecl) string {
	if node.Recv == nil {
		return node.Name.Name
	}

	for _, field := range node.Recv.List {
		if name := recvName(field, node); name != "" {
			return name
		}
	}

	return ""
}

func recvName(field *ast.Field, node *ast.FuncDecl) string {
	switch elem := field.Type.(type) {
	case *ast.StarExpr:
		if name := getX(elem.X); name != "" {
			return name + "_" + node.Name.Name
		}
	case *ast.Ident:
		if unicode.IsUpper([]rune(elem.Name)[0]) {
			return elem.Name + "_" + node.Name.Name
		}
	case *ast.IndexExpr:
		if name := getX(elem.X); name != "" {
			return name + "_" + node.Name.Name
		}
	case *ast.IndexListExpr:
		if name := getX(elem.X); name != "" {
			return name + "_" + node.Name.Name
		}
	}

	return ""
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
