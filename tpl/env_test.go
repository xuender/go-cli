package tpl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/tpl"
)

func TestNewEnv(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	env := tpl.NewEnv()

	ass.NotEqual("", env.Year)
	ass.NotEqual("", env.User)
}

func TestNewEnvByGo(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	env := tpl.NewEnvByGo("env")

	ass.Equal("env.go", env.Path)
	ass.Equal("tpl", env.Package)
}

func TestNewEnvByFile(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	env := tpl.NewEnvByFile("../go", ".mod")

	ass.Equal("../go.mod", env.Path)
}

func TestNewEnvByDir(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	env := tpl.NewEnvByDir(".")

	ass.Equal("tpl", env.Package)
}

func TestGitURL(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("github.com/xuender/go-cli", tpl.GitURL([]byte("url = git@github.com:xuender/go-cli.git")))
}

func TestPackage2url(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("github/xuender/go-cli", tpl.Package2url("github.com/xuender/go-cli"))
	ass.Equal("github", tpl.Package2url("github.com"))
	ass.Equal("github/xuender/go-cli", tpl.Package2url("github/xuender/go-cli"))
}

func TestShortName(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("gh/xuender/go-cli", tpl.ShortName("github.com/xuender/go-cli"))
	ass.Equal("gh", tpl.ShortName("github.com"))
	ass.Equal("gh/xuender/go-cli", tpl.ShortName("github/xuender/go-cli"))
}
