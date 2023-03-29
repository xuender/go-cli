package tpl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/tpl"
)

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("aa_bb", tpl.SnakeCase("aaBb"))
	ass.Equal("aa_bb", tpl.SnakeCase("AaBb"))
}
