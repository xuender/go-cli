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
	names := base.Split(strings.TrimFunc(name, unicode.IsSpace), '-', '_', ' ', '\t')

	for index, str := range names {
		if upper := strings.ToUpper(str); commonInitialisms.Has(upper) {
			names[index] = upper
		} else {
			names[index] = strings.ToUpper(str[0:1]) + str[1:]
		}
	}

	return strings.Join(names, "")
}

// nolint
var commonInitialisms = base.NewSet(
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
)
