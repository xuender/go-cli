package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/cmd"
)

func TestTemplateCmd(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	templateCmd := cmd.TemplateCmd(&cobra.Command{})

	ass.NotNil(templateCmd)
	ass.NotEqual("", templateCmd.Long)
	ass.NotEqual("", templateCmd.Short)
}
