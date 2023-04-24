package utils_test

import (
	"fmt"

	"github.com/xuender/go-cli/utils"
)

// ExampleNewSource is an example function.
func ExampleNewSource() {
	source, err := utils.NewSource("source.go")
	fmt.Println(err)
	fmt.Println(source.Funcs)
	fmt.Println(source.Structs[0].FieldNames)

	// Output:
	// <nil>
	// [Struct_String NewStruct NewSource]
	// [Package Funcs Structs]
}
