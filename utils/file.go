package utils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/xuender/kit/oss"
)

func AppendFile(filename string) *os.File {
	if dir := filepath.Dir(filename); dir != "" && dir != "." && !oss.Exist(dir) {
		lo.Must0(os.MkdirAll(dir, oss.DefaultDirFileMod))
	}

	file := lo.Must1(os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, oss.DefaultFileMode))

	lo.Must1(file.Seek(0, io.SeekEnd))

	return file
}

func CreateFile(filename string) *os.File {
	if dir := filepath.Dir(filename); dir != "" && dir != "." && !oss.Exist(dir) {
		lo.Must0(os.MkdirAll(dir, oss.DefaultDirFileMod))
	}

	file := lo.Must1(os.Create(filename))

	return file
}
