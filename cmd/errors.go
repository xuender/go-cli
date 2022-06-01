package cmd

import "errors"

var ErrEmpty = errors.New(Printer.Sprintf("error empty"))
