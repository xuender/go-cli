package utils

import (
	"regexp"
	"strings"
)

var regCamel = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeCase(camelCase string) string {
	snakeCase := regCamel.ReplaceAllString(camelCase, "${1}_${2}")

	return strings.ToLower(snakeCase)
}
