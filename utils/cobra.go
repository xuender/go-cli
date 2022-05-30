package utils

import (
	"go/ast"
	"strings"
)

// UseCobra 判断是否使用 spf13/cobra.
func UseCobra(root *ast.File) bool {
	for _, imp := range root.Imports {
		if strings.Contains(imp.Path.Value, "spf13/cobra") {
			return true
		}
	}

	return false
}
