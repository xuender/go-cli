package utils_test

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/xuender/go-cli/utils"
)

//go:embed files
var _files embed.FS

// ExampleWalk is an example function.
func ExampleWalk() {
	err := utils.Walk(_files, "files", func(path string, entry fs.DirEntry) error {
		fmt.Println(path, entry.Name())

		return nil
	})

	fmt.Println(err)

	// Output:
	// files/a/aa aaa
	// files/b bb
	// files c
	// <nil>
}
