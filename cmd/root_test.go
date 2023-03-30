package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/cmd"
)

func TestExecute(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	ass.NotPanics(func() {
		cmd.Execute()
	})
}
