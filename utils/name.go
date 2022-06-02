package utils

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/xuender/oils/base"
)

func FileName(name string) string {
	buffer := bytes.Buffer{}

	for index, run := range strings.TrimFunc(name, unicode.IsSpace) {
		if unicode.IsSpace(run) {
			continue
		}

		if unicode.IsUpper(run) {
			if index > 0 {
				buffer.WriteRune('_')
			}

			buffer.WriteRune(unicode.ToLower(run))
		} else {
			buffer.WriteRune(run)
		}
	}

	return buffer.String()
}

func TypeName(name string) string {
	names := base.Split(strings.TrimFunc(name, unicode.IsSpace), base.SepInitialisms, '-', '_', ' ', '\t')

	for index, str := range names {
		if upper := strings.ToUpper(str); base.CommonInitialisms.Has(upper) {
			names[index] = upper
		} else {
			names[index] = strings.ToUpper(str[0:1]) + str[1:]
		}
	}

	return strings.Join(names, "")
}
