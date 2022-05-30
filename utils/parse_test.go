package utils_test

import (
	"testing"

	"github.com/xuender/go-cli/utils"
	"github.com/xuender/oils/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	file, err := utils.Parse("parse.go")

	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, 5, len(file.Imports))

	_, err = utils.Parse("unknown")
	assert.NotNil(t, err)
}

func TestPackageAndFuncs(t *testing.T) {
	t.Parallel()

	pack, funcs := utils.PackageAndFuncs("parse.go")

	assert.Equals(t, []string{"Parse", "PackageAndFuncs"}, funcs)
	assert.Equal(t, "utils", pack)
}

func TestPackageAndFuncs_Struct(t *testing.T) {
	t.Parallel()

	pack, funcs := utils.PackageAndFuncs("data_test.go")

	assert.Equals(t, []string{"PublicStruct_PublicFunc"}, funcs)
	assert.Equal(t, "utils_test", pack)
}
