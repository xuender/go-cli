package version_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/go-cli/version"
)

func TestGetVer(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("1.0.1", version.GetVer([]byte(`x.x.x / 2023

* 1.0.2 33

1.0.1 / 2023

* ff

1.0.0 / 2023

* ff`)))
	ass.Equal("1.0.1", version.GetVer([]byte(`x.x.x / 2023

* 1.0.2 33

V1.0.1 / 22

* ff

v1.0.0 / 202

* ff`)))
}

func TestIncVer(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal("1.0.2", version.IncVer("1.0.1"))
	ass.Equal("2.3.10", version.IncVer("2.3.9"))
	ass.Equal("2.3.1000", version.IncVer("2.3.999"))
}

func TestSetVer(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.True(bytes.HasPrefix(version.SetVer([]byte("x.x.x / 123"), "1.0.1"), []byte("1.0.1")))
}
