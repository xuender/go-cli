package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/utils"
)

type Astruct struct{}

func TestParse(t *testing.T) {
	t.Parallel()

	file, err := utils.Parse("parse.go")

	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, 6, len(file.Imports))

	_, err = utils.Parse("unknown")
	assert.NotNil(t, err)
}

func TestPackageAndFuncs(t *testing.T) {
	t.Parallel()

	pack, funcs := utils.PackageAndFuncs("parse.go")

	assert.Equal(t, []string{"Parse", "PackageAndStructs", "PackageAndFuncs", "UseCobra"}, funcs)
	assert.Equal(t, "utils", pack)
}

func TestPackageAndFuncs_Struct(t *testing.T) {
	t.Parallel()

	pack, funcs := utils.PackageAndFuncs("name_test.go")

	assert.Equal(t, []string{"TestSnakeCase"}, funcs)
	assert.Equal(t, "utils_test", pack)
}

func TestPackageAndStructs(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	pack, structs := utils.PackageAndStructs("parse_test.go")

	ass.Equal("utils_test", pack)
	ass.Equal([]string{"Astruct"}, structs)
}

func TestUseCobra(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	file, _ := utils.Parse("parse_test.go")
	ass.False(utils.UseCobra(file))

	file, _ = utils.Parse("..", "cmd", "root.go")
	ass.True(utils.UseCobra(file))
}

func TestPackageAndFuncs_Struct1(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	pack, funcs := utils.PackageAndFuncs("data_test.go")

	ass.Equal([]string{"PublicStruct_PublicFunc", "PublicStruct_PublicFunc2"}, funcs)
	ass.Equal("utils_test", pack)
}

func TestPackageAndFuncs_Struct2(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	pack, funcs := utils.PackageAndFuncs("data2_test.go")

	ass.Equal([]string{"Slice_Clip", "Slice_Cls"}, funcs)
	ass.Equal("utils_test", pack)
}

func TestPackageAndFuncs_Struct3(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	pack, funcs := utils.PackageAndFuncs("data3_test.go")

	ass.Equal([]string{"Map_Has"}, funcs)
	ass.Equal("utils_test", pack)
}
