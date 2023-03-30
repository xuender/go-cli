package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/utils"
)

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("aa_bb", utils.SnakeCase("aaBb"))
	ass.Equal("aa_bb", utils.SnakeCase("AaBb"))
}
