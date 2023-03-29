package tpl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/tpl"
)

type Astruct struct{}

func TestParse(t *testing.T) {
	t.Parallel()

	file, err := tpl.Parse("parse.go")

	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, 5, len(file.Imports))

	_, err = tpl.Parse("unknown")
	assert.NotNil(t, err)
}

func TestPackageAndFuncs(t *testing.T) {
	t.Parallel()

	pack, funcs := tpl.PackageAndFuncs("parse.go")

	assert.Equal(t, []string{"Parse", "PackageAndFuncs"}, funcs)
	assert.Equal(t, "tpl", pack)
}

func TestPackageAndFuncs_Struct(t *testing.T) {
	t.Parallel()

	pack, funcs := tpl.PackageAndFuncs("name_test.go")

	assert.Equal(t, []string{"TestSnakeCase"}, funcs)
	assert.Equal(t, "tpl_test", pack)
}

func TestPackageAndStructs(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	pack, structs := tpl.PackageAndStructs("parse_test.go")

	ass.Equal("tpl_test", pack)
	ass.Equal([]string{"Astruct"}, structs)
}
