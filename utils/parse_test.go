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
	assert.Equal(t, 7, len(file.Imports))

	_, err = utils.Parse("unknown")
	assert.NotNil(t, err)
}

func TestPackageAndFuncs(t *testing.T) {
	t.Parallel()

	pack, funcs := utils.PackageAndFuncs("parse.go")

	assert.Equals(t, []string{"Parse", "PackageAndFuncs", "GitURL"}, funcs)
	assert.Equal(t, "utils", pack)
}

func TestPackageAndFuncs_Struct(t *testing.T) {
	t.Parallel()

	pack, funcs := utils.PackageAndFuncs("data_test.go")

	assert.Equals(t, []string{"PublicStruct_PublicFunc", "PublicStruct_PublicFunc2"}, funcs)
	assert.Equal(t, "utils_test", pack)
}

func TestGitURL(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "github.com/xuender/go-cli", utils.GitURL([]byte(`[core]
        repositoryformatversion = 0
        filemode = true
        bare = false
        logallrefupdates = true
[remote "origin"]
        url = git@github.com:xuender/go-cli.git
        fetch = +refs/heads/*:refs/remotes/origin/*
[branch "main"]
        remote = origin
        merge = refs/heads/main
[branch "dev"]
        remote = origin
        merge = refs/heads/dev`)))
	assert.Equal(t, "github.com/xuender/go-cli", utils.GitURL([]byte(`url = https://github.com/xuender/go-cli.git`)))
}
