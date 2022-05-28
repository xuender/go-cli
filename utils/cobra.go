package utils

import (
	"go/ast"
	"strings"
)

func UseCobra(root *ast.File) bool {
	for _, imp := range root.Imports {
		if strings.Contains(imp.Path.Value, "spf13/cobra") {
			return true
		}
	}

	return false
}

func A() {}
