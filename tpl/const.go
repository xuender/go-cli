package tpl

import (
	"os"
	"path/filepath"

	"github.com/samber/lo"
)

// nolint: gochecknoglobals
var ConfigPath = filepath.Join(lo.Must1(os.UserHomeDir()), ".config", "go-cli")
