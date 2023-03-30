package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/utils"
)

func TestAppendFile(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Panics(func() {
		utils.AppendFile(filepath.Join(os.TempDir(), "go-cli", "test"))
	})

	file := utils.AppendFile("data_test.go")
	defer file.Close()

	ass.NotNil(file)
}

func TestCreateFile(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Panics(func() {
		utils.CreateFile(os.TempDir())
	})

	file := utils.CreateFile(filepath.Join(os.TempDir(), "go-cli", "test"))

	ass.NotNil(file)
	file.Close()
	os.Remove(file.Name())
}
