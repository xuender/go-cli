package utils_test

import (
	"testing"

	"github.com/xuender/go-scaffold/utils"
	"github.com/xuender/oils/assert"
)

func TestUseCobra(t *testing.T) {
	t.Parallel()

	file, _ := utils.Parse("parse.go")
	assert.False(t, utils.UseCobra(file))

	file, _ = utils.Parse("..", "cmd", "root.go")
	assert.True(t, utils.UseCobra(file))
}
