package utils_test

import (
	"testing"

	"github.com/xuender/go-cli/utils"
	"github.com/xuender/oils/assert"
)

func TestFileName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "aa_bb", utils.FileName("aaBb"))
	assert.Equal(t, "aa_bb", utils.FileName("AaBb"))
	assert.Equal(t, "aa_bb", utils.FileName("Aa Bb"))
}

func TestTypeName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Aa", utils.TypeName("aa"))
	assert.Equal(t, "AaBb", utils.TypeName("aa_bb"))
	assert.Equal(t, "AaBb", utils.TypeName("aa bb"))
	assert.Equal(t, "GetURL", utils.TypeName("get_url"))
}
