// Code generated by go-cli. DO NOT EDIT.

package {{ .Package }}

import "github.com/xuender/kit/oss"

// nolint: gochecknoinits
func init() {
	if oss.Version == "" {
		oss.Version = "{{ .Version }}"
	}

	if oss.BuildTime == "" || oss.BuildTime[0] == '0' {
		oss.BuildTime = "{{ .BuildTime }}"
	}
}
